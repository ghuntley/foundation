// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package golang

import (
	"context"
	"fmt"
	"strings"

	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/util/system"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"namespacelabs.dev/foundation/build"
	"namespacelabs.dev/foundation/build/buildkit"
	"namespacelabs.dev/foundation/internal/artifacts/oci"
	"namespacelabs.dev/foundation/internal/engine/ops"
	"namespacelabs.dev/foundation/internal/llbutil"
	"namespacelabs.dev/foundation/internal/production"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/workspace/compute"
	"namespacelabs.dev/foundation/workspace/pins"
	"namespacelabs.dev/foundation/workspace/tasks"
)

var useSeparateGoModPhase = false

func buildUsingBuildkit(ctx context.Context, env ops.Environment, bin GoBinary, conf build.Configuration) (compute.Computable[oci.Image], error) {
	local := buildkit.LocalContents{
		Module:         conf.Workspace,
		Path:           bin.GoModulePath,
		ObserveChanges: bin.isFocus,
	}

	src := buildkit.MakeLocalState(local)

	base := makeGoBuildBase(bin.GoVersion, buildkit.HostPlatform())

	prodBase, err := production.ServerImageLLB(production.Distroless, *conf.Target)
	if err != nil {
		return nil, err
	}

	label := "building"
	if bin.PackageName != "" {
		label += fmt.Sprintf(" %s", bin.PackageName)
	}

	goBuild := []string{"go", "build", "-v"}

	if env.Proto().GetPurpose() == schema.Environment_PRODUCTION {
		goBuild = append(goBuild, "-trimpath")
	}

	goBuild = append(goBuild, "-o", "/out/%s")

	state := (llbutil.RunGo{
		Base:       prepareGoMod(base, src, conf.Target).Root(),
		SrcMount:   src,
		WorkingDir: bin.SourcePath,
		Platform:   conf.Target,
	}).With(
		llbutil.PrefixSh(label, conf.Target, strings.Join(goBuild, " "), bin.BinaryName)...).
		AddMount("/out", prodBase)

	image, err := buildkit.LLBToImage(ctx, env, conf.Target, state, local)
	if err != nil {
		return nil, err
	}

	return compute.Named(
		tasks.Action("go.build.binary").WellKnown(tasks.WkModule, bin.ModuleName), image), nil
}

func prepareGoMod(base, src llb.State, platform *specs.Platform) llb.ExecState {
	r := llbutil.RunGo{
		Base:       base,
		SrcMount:   src,
		WorkingDir: ".",
		Platform:   platform,
	}

	ro := llbutil.PrefixSh("updating deps", platform, "go mod download -x")
	if !useSeparateGoModPhase {
		return r.With(ro...)
	}

	return r.PrepareGoMod(ro...)
}

func makeGoImage(version string) string {
	return fmt.Sprintf("docker.io/library/golang:%s-alpine", version)
}

func goAlpine(version string, platform specs.Platform) llb.State {
	img := makeGoImage(version)

	if r, err := pins.CheckImage(img); err == nil {
		return llbutil.Image(r, platform)
	}

	return llbutil.Image(img, platform)
}

func makeGoBuildBase(version string, platform specs.Platform) llb.State {
	return goAlpine(version, platform).
		AddEnv("CGO_ENABLED", "0").
		AddEnv("PATH", "/usr/local/go/bin:"+system.DefaultPathEnvUnix).
		AddEnv("GOPATH", "/go").
		Run(llb.Shlex("apk add --no-cache git"),
			llb.WithCustomName("[prepare build image] apk add --no-cache git")).Root()
}