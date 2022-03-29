// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package build

import (
	"context"
	"strings"
	"time"

	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"namespacelabs.dev/foundation/internal/artifacts/oci"
	"namespacelabs.dev/foundation/internal/engine/ops"
	"namespacelabs.dev/foundation/internal/wscontents"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/workspace/compute"
	"namespacelabs.dev/foundation/workspace/devhost"
)

var (
	FixedPoint       = time.Unix(1, 1)
	platformOverride = []specs.Platform{}
)

type Spec interface {
	BuildImage(context.Context, ops.Environment, Configuration) (compute.Computable[oci.Image], error)
	PlatformIndependent() bool
}

type Plan struct {
	SourcePackage schema.PackageName
	SourceLabel   string
	Spec          Spec
	Workspace     Workspace
	Platforms     []specs.Platform
}

type Workspace interface {
	ModuleName() string
	Abs() string
	VersionedFS(rel string, observeChanges bool) compute.Computable[wscontents.Versioned]
}

type Configuration struct {
	SourceLabel string
	Target      *specs.Platform
	Workspace   Workspace
}

type BuildPlatformsVar struct{}

func (BuildPlatformsVar) String() string {
	var p []string
	for _, plat := range platformOverride {
		p = append(p, devhost.FormatPlatform(plat))
	}
	return strings.Join(p, ",")
}

func (BuildPlatformsVar) Set(s string) error {
	platformParts := strings.Split(s, ",")

	var ps []specs.Platform
	for _, p := range platformParts {
		parsed, err := devhost.ParsePlatform(p)
		if err != nil {
			return err
		}
		ps = append(ps, parsed)
	}

	platformOverride = ps
	return nil
}

func (BuildPlatformsVar) Type() string {
	return ""
}

func PlatformsOrOverrides(platforms []specs.Platform) []specs.Platform {
	if len(platformOverride) > 0 {
		return platformOverride
	}
	return platforms
}