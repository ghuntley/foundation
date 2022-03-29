// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package golang

import (
	"context"
	"path/filepath"

	"namespacelabs.dev/foundation/build"
	"namespacelabs.dev/foundation/internal/artifacts/oci"
	"namespacelabs.dev/foundation/internal/engine/ops"
	"namespacelabs.dev/foundation/internal/gosupport"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/workspace"
	"namespacelabs.dev/foundation/workspace/compute"
)

type GoBinary struct {
	PackageName schema.PackageName `json:"packageName"`
	ModuleName  string             `json:"moduleName"`

	GoModulePath string   `json:"modulePath"` // Relative to workspace root.
	GoModule     string   `json:"module"`     // Go module name.
	GoVersion    string   `json:"goVersion"`
	SourcePath   string   `json:"sourcePath"` // Relative to workspace root.
	BinaryName   string   `json:"binaryName"`
	Capabilities []string `json:"capabilities"`

	UnsafeCacheable bool // Unsafe because we can't guarantee that the sources used for compilation are consistent with the workspace contents.

	isFocus bool
}

var UseBuildKitForBuilding = false

func (gb GoBinary) BuildImage(ctx context.Context, env ops.Environment, conf build.Configuration) (compute.Computable[oci.Image], error) {
	if UseBuildKitForBuilding {
		return buildUsingBuildkit(ctx, env, gb, conf)
	}

	return Build(ctx, env, gb, conf)
}

func (gb GoBinary) PlatformIndependent() bool { return false }

func FromLocation(loc workspace.Location, pkgName string) (*GoBinary, error) {
	absSrc := loc.Abs(pkgName)
	mod, modFile, err := gosupport.LookupGoModule(absSrc)
	if err != nil {
		return nil, err
	}

	relMod, err := filepath.Rel(loc.Module.Abs(), modFile)
	if err != nil {
		return nil, err
	}

	return &GoBinary{
		PackageName:  loc.PackageName,
		ModuleName:   loc.Module.ModuleName(),
		GoModulePath: filepath.Dir(relMod),
		GoModule:     mod.Module.Mod.Path,
		SourcePath:   loc.Rel(pkgName),
		GoVersion:    mod.Go.Version,
	}, nil
}