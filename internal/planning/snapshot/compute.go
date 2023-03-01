// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package snapshot

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"namespacelabs.dev/foundation/framework/rpcerrors/multierr"
	"namespacelabs.dev/foundation/internal/bytestream"
	"namespacelabs.dev/foundation/internal/compute"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/filewatcher"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/fnfs"
	"namespacelabs.dev/foundation/internal/parsing"
	"namespacelabs.dev/foundation/internal/planning"
	"namespacelabs.dev/foundation/internal/planning/eval"
	"namespacelabs.dev/foundation/internal/runtime"
	"namespacelabs.dev/foundation/internal/wscontents"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/std/cfg"
	"namespacelabs.dev/foundation/std/pkggraph"

	"namespacelabs.dev/foundation/std/tasks"
)

func RequireServers(env cfg.Context, planner runtime.Planner, servers ...schema.PackageName) compute.Computable[*ServerSnapshot] {
	return &requiredServers{env: env, planner: planner, packages: servers}
}

type requiredServers struct {
	env      cfg.Context
	planner  runtime.Planner
	packages []schema.PackageName

	compute.LocalScoped[*ServerSnapshot]
}

type ServerSnapshot struct {
	stack  *planning.Stack
	sealed pkggraph.SealedPackageLoader
	// Used in Observe()
	env      cfg.Context
	planner  runtime.Planner
	packages []schema.PackageName
}

var _ compute.Versioned = &ServerSnapshot{}

func (rs *requiredServers) Action() *tasks.ActionEvent {
	return tasks.Action("planning.require-servers")
}

func (rs *requiredServers) Inputs() *compute.In {
	return compute.Inputs().Indigestible("env", rs.env).Strs("packages", schema.Strs(rs.packages...))
}

func (rs *requiredServers) Output() compute.Output {
	return compute.Output{NotCacheable: true}
}

func (rs *requiredServers) Compute(ctx context.Context, _ compute.Resolved) (*ServerSnapshot, error) {
	return computeSnapshot(ctx, rs.env, rs.planner, rs.packages)
}

func computeSnapshot(ctx context.Context, env cfg.Context, planner runtime.Planner, packages []schema.PackageName) (*ServerSnapshot, error) {
	pl := parsing.NewPackageLoader(env)

	var servers []planning.Server
	for _, pkg := range packages {
		server, err := planning.RequireServerWith(ctx, env, pl, schema.PackageName(pkg))
		if err != nil {
			return nil, err
		}

		servers = append(servers, server)
	}

	stack, err := planning.ComputeStack(ctx, servers, planning.ProvisionOpts{Planner: planner, PortRange: eval.DefaultPortRange()})
	if err != nil {
		return nil, err
	}

	return &ServerSnapshot{stack: stack, sealed: pl.Seal(), env: env, planner: planner, packages: packages}, nil
}

func (snap *ServerSnapshot) Get(pkgs ...schema.PackageName) ([]planning.PlannedServer, error) {
	var servers []planning.PlannedServer
	for _, pkg := range pkgs {
		srv, ok := snap.stack.Get(pkg)
		if !ok {
			return nil, fnerrors.InternalError("%s: not present in the snapshot", pkg)
		}
		servers = append(servers, srv)
	}
	return servers, nil
}

func (snap *ServerSnapshot) Modules() pkggraph.Modules {
	return snap.sealed
}

func (snap *ServerSnapshot) Env() pkggraph.Context {
	return pkggraph.MakeSealedContext(snap.env, snap.sealed)
}

func (snap *ServerSnapshot) Equals(rhs *ServerSnapshot) bool {
	return false // XXX optimization.
}

func (snap *ServerSnapshot) Stack() *planning.Stack {
	return snap.stack
}

func (snap *ServerSnapshot) Observe(ctx context.Context, onChange func(compute.ResultWithTimestamp[any], compute.ObserveNote)) (func(), error) {
	obs := obsState{onChange: onChange}
	if err := obs.prepare(ctx, snap); err != nil {
		return nil, err
	}
	return obs.cancel, nil
}

type obsState struct {
	cancelWatcher func()
	onChange      func(compute.ResultWithTimestamp[any], compute.ObserveNote)
}

func (p *obsState) prepare(ctx context.Context, snapshot *ServerSnapshot) error {
	cancel, err := observe(ctx, snapshot, func(newSnapshot *ServerSnapshot) {
		if p.cancelWatcher != nil {
			p.cancelWatcher()
			p.cancelWatcher = nil
		}

		r := compute.ResultWithTimestamp[any]{
			Completed: time.Now(),
		}
		r.Value = newSnapshot

		p.onChange(r, compute.ObserveContinuing)

		if err := p.prepare(ctx, newSnapshot); err != nil {
			compute.Stop(ctx, err)
		}
	})
	p.cancelWatcher = cancel
	return err
}

func (p *obsState) cancel() {
	if p.cancelWatcher != nil {
		p.cancelWatcher()
	}
}

func observe(ctx context.Context, snap *ServerSnapshot, onChange func(*ServerSnapshot)) (func(), error) {
	logger := console.TypedOutput(ctx, "observepackages", console.CatOutputUs)

	watcher, err := filewatcher.NewFactory(ctx)
	if err != nil {
		return nil, err
	}

	merged := map[string]*pkggraph.Package{}

	var fileCount, packageCount int
	for _, pkg := range snap.sealed.Packages() {
		// Don't monitor changes to external modules, assume they're immutable.
		if pkg.Location.Module.IsExternal() {
			continue
		}

		packageCount++

		if err := fnfs.VisitFiles(ctx, pkg.PackageSources, func(path string, _ bytestream.ByteStream, _ fs.DirEntry) error {
			fileCount++
			abs := filepath.Join(pkg.Location.Module.Abs(), path)
			merged[abs] = pkg
			return watcher.AddFile(abs) // Path is relative to module root.
		}); err != nil {
			watcher.Close()
			return nil, fnerrors.InternalError("%s: failed to visit sources: %w", pkg.Location.PackageName, err)
		}
	}

	fmt.Fprintf(console.Debug(ctx), "observing pkggraph: merged view has %d files over %d packages\n", fileCount, packageCount)

	bufferCh := make(chan []fsnotify.Event)
	go func() {
		for buffer := range bufferCh {
			var dirty schema.PackageList
			var errs []error
			for _, ev := range buffer {
				pkg := merged[ev.Name]
				if pkg == nil {
					continue
				}

				relPath, err := filepath.Rel(pkg.Location.Module.Abs(), ev.Name)
				if err != nil {
					fmt.Fprintf(console.Debug(ctx), "failed to calculate relative path of %s: %v\n", ev.Name, err)
					continue
				}

				contents, contentsErr := os.ReadFile(ev.Name)
				expected, expectedErr := fs.ReadFile(pkg.PackageSources, relPath)
				if contentsErr == nil && expectedErr == nil && bytes.Equal(contents, expected) {
					continue
				}

				dirty.Add(pkg.PackageName())
				errs = append(errs, contentsErr)
				errs = append(errs, expectedErr)
			}

			if dirty.Len() == 0 {
				continue
			}

			if err := multierr.New(errs...); err != nil {
				fmt.Fprintf(console.Warnings(ctx), "Got errors while watching for changes:\n  %v\n", err)
			}

			newSnapshot, err := computeSnapshot(ctx, snap.env, snap.planner, snap.packages)
			if err != nil {
				if msg, ok := fnerrors.IsExpected(err); ok {
					fmt.Fprintf(console.Stderr(ctx), "\n  %s\n\n", msg)
					continue // Swallow the error in case it is expected.
				}
				compute.Stop(ctx, err)
				break
			}

			if newSnapshot.Equals(snap) {
				continue
			}

			fmt.Fprintf(logger, "Recomputed graph, due to changes to %s.\n", strings.Join(dirty.PackageNamesAsString(), ", "))

			onChange(newSnapshot)
			break // Don't send any more events.
		}

		for range bufferCh {
			// Drain the channel in case we escaped the loop above, so the go-routine below
			// has a chance to observe a canceled context and close the channel.
		}
	}()

	w, err := watcher.StartWatching(ctx)
	if err != nil {
		return nil, err
	}

	go func() {
		wscontents.AggregateFSEvents(w, console.Debug(ctx), logger, bufferCh)
	}()

	return func() {
		w.Close()
	}, nil
}
