// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package golang

import (
	"context"
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/util/system"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"namespacelabs.dev/foundation/build"
	"namespacelabs.dev/foundation/build/buildkit"
	"namespacelabs.dev/foundation/internal/artifacts/oci"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/engine/ops"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/gosupport"
	"namespacelabs.dev/foundation/internal/llbutil"
	"namespacelabs.dev/foundation/internal/production"
	"namespacelabs.dev/foundation/provision"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/workspace"
	"namespacelabs.dev/foundation/workspace/compute"
	"namespacelabs.dev/foundation/workspace/pins"
	"namespacelabs.dev/foundation/workspace/tasks"
)

var (
	useSeparateGoModPhase = false
)

type buildConf interface {
	build.BuildTarget
	build.BuildWorkspace
}

func buildUsingBuildkit(ctx context.Context, env ops.Environment, bin GoBinary, conf buildConf) (compute.Computable[oci.Image], error) {
	local := buildkit.LocalContents{
		Module:         conf.Workspace(),
		Path:           bin.GoModulePath,
		ObserveChanges: bin.isFocus,
	}

	src := buildkit.MakeLocalState(local)

	base := makeGoBuildBase(ctx, bin.GoVersion, buildkit.HostPlatform())

	var prodBase llb.State

	if !bin.BinaryOnly {
		var err error
		prodBase, err = production.ServerImageLLB(production.Distroless, *conf.TargetPlatform())
		if err != nil {
			return nil, err
		}
	} else {
		prodBase = llb.Scratch()
	}

	label := "building"
	if bin.PackageName != "" {
		label += fmt.Sprintf(" %s", bin.PackageName)
	}

	goBuild := goBuildArgs(bin.GoVersion)
	goBuild = append(goBuild, fmt.Sprintf("-o=/out/%s", bin.BinaryName))

	state := (llbutil.RunGo{
		Base:       prepareGoMod(base, src, conf.TargetPlatform()).Root(),
		SrcMount:   src,
		WorkingDir: bin.SourcePath,
		Platform:   conf.TargetPlatform(),
	}).With(
		llbutil.PrefixSh(label, conf.TargetPlatform(), "go "+strings.Join(goBuild, " "))...).
		AddMount("/out", prodBase)

	image, err := buildkit.LLBToImage(ctx, env, conf, state, local)
	if err != nil {
		return nil, err
	}

	return compute.Named(
		tasks.Action("go.build.binary").Scope(bin.PackageName).WellKnown(tasks.WkModule, bin.ModuleName), image), nil
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

func goAlpine(ctx context.Context, version string, platform specs.Platform) llb.State {
	img := makeGoImage(version)

	if r, err := pins.CheckImage(img); err == nil {
		return llbutil.Image(r, platform)
	}

	fmt.Fprintf(console.Warnings(ctx), "go: no pinned version of %q\n", img)

	return llbutil.Image(img, platform)
}

func makeGoBuildBase(ctx context.Context, version string, platform specs.Platform) llb.State {
	st := goAlpine(ctx, version, platform).
		AddEnv("CGO_ENABLED", "0").
		AddEnv("PATH", "/usr/local/go/bin:"+system.DefaultPathEnvUnix).
		AddEnv("GOPATH", "/go").
		Run(llb.Shlex("apk add --no-cache git"),
			llb.WithCustomName("[prepare build image] apk add --no-cache git")).Root()

	if llbutil.GitCredentialsBuildkitSecret != "" {
		st = st.Run(llb.Shlex("git config --global credential.helper store")).Root()
	}

	return st
}

func tidyBuildkit(ctx context.Context, env provision.Env, loc workspace.Location, server *schema.Server) error {
	_, modPath, err := gosupport.LookupGoModule(loc.Abs())
	if err != nil {
		return err
	}

	local := buildkit.LocalContents{
		Module:         loc.Module,
		Path:           modPath,
		ObserveChanges: false,
	}

	p := buildkit.HostPlatform()

	src := buildkit.MakeLocalState(local)

	ext := &FrameworkExt{}
	if err := workspace.MustExtension(server.Ext, ext); err != nil {
		return fnerrors.Wrap(loc, err)
	}
	base := makeGoBuildBase(ctx, ext.GoVersion, p)

	state := (llbutil.RunGo{
		Base:       base,
		SrcMount:   src,
		WorkingDir: ".", // Correct?
		Platform:   &p,
	}).With(
		llbutil.PrefixSh("go mod tidy", &p, "go mod tidy")...)

	output, err := buildkit.LLBToFS(ctx, env,
		build.NewBuildTarget(&p).WithWorkspace(loc.Module), state.Root())
	if err != nil {
		return err
	}

	fsys, err := compute.GetValue(ctx, output)
	if err != nil {
		return err
	}

	for _, file := range []string{"go.mod", "go.sum"} {
		data, err := fs.ReadFile(fsys, file)
		if err != nil {
			return err
		}

		ioutil.WriteFile(filepath.Join(filepath.Dir(modPath), file), data, 0644)
	}

	return nil
}
