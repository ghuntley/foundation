// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package wsremote

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/moby/patternmatcher"
	"namespacelabs.dev/foundation/internal/bytestream"
	"namespacelabs.dev/foundation/internal/compute"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/fnfs/digestfs"
	"namespacelabs.dev/foundation/internal/fnfs/memfs"
	"namespacelabs.dev/foundation/internal/wscontents"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/std/tasks"
)

type Sink interface {
	Deposit(context.Context, []*wscontents.FileEvent) (bool, error)
}

// Returns a wscontents.Versioned which will produce a local snapshot as expected
// but forwards all filesystem events (e.g. changes, removals) to the specified sink.
func ObserveAndPush(absPath string, excludes []string, sink Sink, digestMode bool, extraInputs ...compute.UntypedComputable) compute.Computable[any] {
	return &observePath{absPath: absPath, excludes: excludes, sink: sink, digestMode: digestMode, extraInputs: extraInputs}
}

type observePath struct {
	absPath     string
	excludes    []string
	sink        Sink
	digestMode  bool
	extraInputs []compute.UntypedComputable

	compute.LocalScoped[any]
}

func (op *observePath) Action() *tasks.ActionEvent {
	return tasks.Action("web.contents.observe")
}

func (op *observePath) Inputs() *compute.In {
	in := compute.Inputs().Str("absPath", op.absPath).Indigestible("not cacheable", "true")
	for k, extra := range op.extraInputs {
		in = in.Computable(fmt.Sprintf("extra:%d", k), extra)
	}
	return in
}

func (op *observePath) Compute(ctx context.Context, _ compute.Resolved) (any, error) {
	fmt.Fprintf(console.Debug(ctx), "wsremote: starting w/ snapshotting %q (excludes: %v)\n", op.absPath, op.excludes)

	excludeMatcher, err := patternmatcher.New(op.excludes)
	if err != nil {
		return nil, err
	}

	snapshot, err := wscontents.SnapshotDirectory(ctx, op.absPath, excludeMatcher, op.digestMode)
	if err != nil {
		return nil, err
	}

	return localObserver{absPath: op.absPath, excludeMatcher: excludeMatcher, digestMode: op.digestMode, snapshot: snapshot, sink: op.sink}, nil
}

type localObserver struct {
	absPath        string
	excludeMatcher *patternmatcher.PatternMatcher
	digestMode     bool
	snapshot       *memfs.FS
	sink           Sink
}

func (lo localObserver) Abs() string { return lo.absPath }
func (lo localObserver) FS() fs.FS   { return lo.snapshot }
func (lo localObserver) ComputeDigest(ctx context.Context) (schema.Digest, error) {
	return digestfs.Digest(ctx, lo.snapshot)
}

func (lo localObserver) Observe(ctx context.Context, onChange func(compute.ResultWithTimestamp[any], compute.ObserveNote)) (func(), error) {
	// XXX we're doing polling for correctness; this needs to use filewatcher.

	// This observer is special; if we know that the scheduler wants to observe
	// the graph, then we trigger a syncing of local files to the destination
	// sink. We don't actually ever emit a new snapshot.

	closeCh := make(chan struct{})
	last := lo.snapshot

	go func() {
		t := time.NewTicker(time.Second)
		defer t.Stop()

		for {
			select {
			case <-closeCh:
				return
			case <-t.C:
				newSnapshot, deposited, err := checkSnapshot(ctx, last, lo.absPath, lo.excludeMatcher, lo.digestMode, lo.sink)
				if err != nil {
					fmt.Fprintf(console.Errors(ctx), "FileSync failed while snapshotting %q: %v\n", lo.absPath, err)
					return
				}

				if !deposited {
					r := compute.ResultWithTimestamp[any]{
						Completed: time.Now(),
					}
					r.Value = localObserver{absPath: lo.absPath, snapshot: newSnapshot, sink: lo.sink}
					onChange(r, compute.ObserveContinuing)
				}

				last = newSnapshot
			}
		}
	}()

	return func() { close(closeCh) }, nil
}

func checkSnapshot(ctx context.Context, previous *memfs.FS, absPath string, excludeMatcher *patternmatcher.PatternMatcher, digestMode bool, sink Sink) (*memfs.FS, bool, error) {
	newSnapshot, err := wscontents.SnapshotDirectory(ctx, absPath, excludeMatcher, digestMode)
	if err != nil {
		return nil, false, err
	}

	// First we iterate over all files in the new snapshot. This index will be
	// used to signal which files are actually new, and which files were removed
	// from the old snapshot.
	newFiles := map[string]bytestream.Static{}
	newFilesModes := map[string]memfs.FileDirent{}
	if err := newSnapshot.VisitFilesWithoutContext(func(path string, bs bytestream.Static, de memfs.FileDirent) error {
		newFiles[path] = bs
		newFilesModes[path] = de
		return nil
	}); err != nil {
		return nil, false, err
	}

	var events []*wscontents.FileEvent

	if err := previous.VisitFilesWithoutContext(func(path string, bs bytestream.Static, _ memfs.FileDirent) error {
		if newFile, ok := newFiles[path]; !ok {
			events = append(events, &wscontents.FileEvent{Path: path, Event: wscontents.FileEvent_REMOVE})
			delete(newFiles, path)
		} else {
			if bytes.Equal(bs.Contents, newFile.Contents) {
				// No changes, don't re-write.
				delete(newFiles, path)
			}
		}
		return nil
	}); err != nil {
		return nil, false, err
	}

	m := checkMkdir{previous: previous, newdirs: map[string]struct{}{}}

	for filename, contents := range newFiles {
		if err := m.check(filepath.Dir(filename)); err != nil {
			return nil, false, err
		}

		ev := &wscontents.FileEvent{
			Event: wscontents.FileEvent_WRITE,
			Path:  filename,
			Mode:  uint32(newFilesModes[filename].FileMode.Perm()),
		}

		if !digestMode {
			ev.NewContents = contents.Contents
		}

		events = append(events, ev)
	}

	// Mkdirs come first.
	events = append(m.events, events...)

	var deposited bool
	var depositErr error

	if len(events) > 0 {
		deposited, depositErr = sink.Deposit(ctx, events)
	} else {
		// If there are no changes consider everything done
		// and avoid generating a new version in Observe() for nothing.
		deposited = true
	}

	return newSnapshot, deposited, depositErr
}

type checkMkdir struct {
	previous *memfs.FS
	newdirs  map[string]struct{}
	events   []*wscontents.FileEvent
}

// XXX this doesn't handle the case where a directory becomes a file, or vice-versa.
func (m *checkMkdir) check(dir string) error {
	if dir == "." {
		return nil
	}

	if err := m.check(filepath.Dir(dir)); err != nil {
		return err
	}

	if _, ok := m.newdirs[dir]; ok {
		return nil
	}

	if _, err := m.previous.Open(dir); err != nil {
		if !os.IsNotExist(err) {
			return err
		}

		m.events = append(m.events, &wscontents.FileEvent{Path: dir, Event: wscontents.FileEvent_MKDIR})
		m.newdirs[dir] = struct{}{}
	}

	return nil
}
