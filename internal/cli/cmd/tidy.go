// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package cmd

import (
	"context"
	"io"
	"sort"
	"strings"
	"sync"

	"github.com/spf13/cobra"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/fnfs"
	"namespacelabs.dev/foundation/internal/fnfs/memfs"
	"namespacelabs.dev/foundation/internal/frontend/fncue"
	"namespacelabs.dev/foundation/languages"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/workspace"
	"namespacelabs.dev/foundation/workspace/module"
	"tailscale.com/util/multierr"
)

func NewTidyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tidy",
		Short: "Ensures that each server has the appropriate dependencies configured.",

		RunE: fncobra.RunE(func(ctx context.Context, args []string) error {
			// First of all, we work through all packages to make sure we have captured
			// their dependencies locally. If we don't do this here, package parsing below
			// will fail.

			root, err := module.FindRoot(ctx, ".")
			if err != nil {
				return err
			}

			pl := workspace.NewPackageLoader(root)

			if err := fillDependencies(ctx, root, pl); err != nil {
				return err
			}

			list, err := workspace.ListSchemas(ctx, root)
			if err != nil {
				return err
			}

			packages := []*workspace.Package{}
			for _, loc := range list.Locations {
				pkg, err := pl.LoadByName(ctx, loc.AsPackageName())
				if err != nil {
					return err
				}

				if pkg.Binary != nil {
					continue
				}

				packages = append(packages, pkg)
			}

			var errs []error
			for _, pkg := range packages {
				switch {
				case pkg.Server != nil:
					lang := languages.IntegrationFor(pkg.Server.Framework)
					if err := lang.TidyServer(ctx, pl, pkg.Location, pkg.Server); err != nil {
						errs = append(errs, err)
					}

				case pkg.Node() != nil:
					for _, fmwk := range pkg.Node().CodegeneratedFrameworks() {
						lang := languages.IntegrationFor(fmwk)
						if err := lang.TidyNode(ctx, pl, pkg); err != nil {
							errs = append(errs, err)
						}
					}
				}
			}
			for _, fmwk := range schema.Framework_value {
				lang := languages.IntegrationFor(schema.Framework(fmwk))
				if lang == nil {
					continue
				}
				if err := lang.TidyWorkspace(ctx, packages); err != nil {
					errs = append(errs, err)
				}
			}

			return multierr.New(errs...)
		}),
	}

	return cmd
}

func fillDependencies(ctx context.Context, root *workspace.Root, pl *workspace.PackageLoader) error {
	list, err := workspace.ListSchemas(ctx, root)
	if err != nil {
		return err
	}

	alloc := &allocator{
		loader:   pl,
		root:     root,
		resolved: map[string]*schema.Workspace_Dependency{},
		modules:  map[string]*schema.Workspace_Dependency{},
		left:     append([]fnfs.Location{}, list.Locations...),
	}

	for {
		alloc.mu.Lock()
		var loc *fnfs.Location
		if len(alloc.left) > 0 {
			loc = &alloc.left[0]
			alloc.left = alloc.left[1:]
		}
		alloc.mu.Unlock()

		if loc == nil {
			break
		}

		r := &workspaceLoader{alloc}
		imports := map[string]*fncue.CuePackage{}

		// Check whether imports refer to packages; we'll see calls to workspaceResolver.
		// We ignore errors, because some of the errors may be related to the lack of
		// presence of packages.
		_ = fncue.CollectImports(ctx, r, loc.AsPackageName().String(), imports)

		parsed, err := alloc.loader.LoadByNameWithOpts(ctx, loc.AsPackageName(), workspace.DontLoadDependencies())
		if err != nil {
			return err
		}

		switch {
		case parsed.Server != nil:
			if err := alloc.checkResolves(ctx, parsed.Server.Import, parsed.Server.Reference); err != nil {
				return err
			}
		case parsed.Service != nil, parsed.Extension != nil:
			if err := alloc.checkResolves(ctx, parsed.Node().Import, parsed.Node().Reference); err != nil {
				return err
			}
		}
	}

	root.Workspace.Dep = nil

	modules := map[string]*schema.Workspace_Dependency{}
	for _, dep := range alloc.resolved {
		if modules[dep.ModuleName] != nil {
			continue
		}

		modules[dep.ModuleName] = dep
		root.Workspace.Dep = append(root.Workspace.Dep, dep)
	}

	return rewriteWorkspace(ctx, root, root.Workspace)
}

func rewriteWorkspace(ctx context.Context, root *workspace.Root, ws *schema.Workspace) error {
	sort.Slice(ws.Dep, func(i, j int) bool {
		return strings.Compare(ws.Dep[i].ModuleName, ws.Dep[j].ModuleName) < 0
	})

	// Write an updated workspace.ns.textpb before continuing.
	return fnfs.WriteWorkspaceFile(ctx, console.Stdout(ctx), root.FS(), workspace.WorkspaceFilename, func(w io.Writer) error {
		return workspace.FormatWorkspace(w, ws)
	})
}

type allocator struct {
	loader   *workspace.PackageLoader
	root     *workspace.Root
	mu       sync.Mutex // Protects resolved and left.
	resolved map[string]*schema.Workspace_Dependency
	modules  map[string]*schema.Workspace_Dependency
	left     []fnfs.Location
}

func (alloc *allocator) checkResolves(ctx context.Context, pkgs []string, refs []*schema.Reference) error {
	for _, pkg := range pkgs {
		if _, err := alloc.checkResolve(ctx, schema.PackageName(pkg)); err != nil {
			return err
		}
	}

	for _, ref := range refs {
		if ref.PackageName == "" {
			continue
		}
		if _, err := alloc.checkResolve(ctx, schema.PackageName(ref.PackageName)); err != nil {
			return err
		}
	}

	return nil
}

func (alloc *allocator) checkResolve(ctx context.Context, sch schema.PackageName) (workspace.Location, error) {
	if _, ok := schema.IsParent(alloc.root.Workspace.ModuleName, sch); ok {
		return alloc.loader.Resolve(ctx, sch)
	}

	// Check if we already processed this package.
	alloc.mu.Lock()
	resolved := alloc.resolved[sch.String()]
	alloc.mu.Unlock()

	var didResolve bool
	if resolved == nil {
		// First, is there a replace statement that applies to this package?
		replaced, err := alloc.loader.MatchModuleReplace(sch)
		if err != nil {
			return workspace.Location{}, err
		}

		// If so, there's nothing for us to do here.
		if replaced != nil {
			return alloc.loader.Resolve(ctx, sch)
		}

		// Then, resolve the package to a module name + relative path.
		mod, err := workspace.ResolveModule(ctx, sch.String())
		if err != nil {
			return workspace.Location{}, err
		}

		// Check if we already parsed this module.
		alloc.mu.Lock()
		resolved = alloc.modules[mod.ModuleName]
		alloc.mu.Unlock()

		// No? Then fetch the latest head.
		if resolved == nil {
			dep, err := workspace.ModuleHead(ctx, mod)
			if err != nil {
				return workspace.Location{}, err
			}
			resolved = dep
		}

		didResolve = true
	}

	loc, err := alloc.loader.ExternalLocation(ctx, resolved, sch)
	if err == nil && didResolve {
		alloc.mu.Lock()
		alloc.resolved[sch.String()] = resolved
		alloc.modules[resolved.ModuleName] = resolved
		// If we just parsed this package, add it to the queue of packages to
		// be checked for references as well.
		alloc.left = append(alloc.left, fnfs.Location{
			ModuleName: loc.Module.ModuleName(),
			RelPath:    loc.Rel(),
		})
		alloc.mu.Unlock()
	}

	return loc, err
}

type workspaceLoader struct {
	alloc *allocator
}

func (wr *workspaceLoader) SnapshotDir(ctx context.Context, sch schema.PackageName, opts memfs.SnapshotOpts) (fnfs.Location, error) {
	loc, err := wr.alloc.checkResolve(ctx, sch)
	if err != nil {
		return fnfs.Location{}, err
	}

	fsys, err := memfs.SnapshotDir(fnfs.Local(loc.Module.Abs()), loc.Rel(), opts)
	if err != nil {
		return fnfs.Location{}, err
	}

	return fnfs.Location{
		ModuleName: loc.Module.ModuleName(),
		RelPath:    loc.Rel(),
		FS:         fsys,
	}, nil
}
