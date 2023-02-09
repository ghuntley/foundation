// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package golang

import (
	"context"
	"fmt"
	"strings"

	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/util/system"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"namespacelabs.dev/foundation/internal/artifacts/oci"
	"namespacelabs.dev/foundation/internal/build"
	"namespacelabs.dev/foundation/internal/build/buildkit"
	"namespacelabs.dev/foundation/internal/compute"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/dependencies/pins"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/git"
	"namespacelabs.dev/foundation/internal/llbutil"
	"namespacelabs.dev/foundation/std/pkggraph"
)

var (
	useSeparateGoModPhase = false
)

func buildUsingBuildkit(ctx context.Context, env pkggraph.SealedContext, bin GoBinary, conf build.Configuration) (compute.Computable[oci.Image], error) {
	local := buildkit.LocalContents{
		Module: conf.Workspace(),
		Path:   bin.GoModulePath,
	}

	src := buildkit.MakeLocalState(local)

	if conf.TargetPlatform() == nil {
		return nil, fnerrors.InternalError("go: target platform is missing")
	}

	base := makeGoBuildBase(ctx, bin.GoVersion, *conf.TargetPlatform())

	var prodBase llb.State

	if !bin.BinaryOnly {
		prodBase0, err := baseProdImage(ctx, env, *conf.TargetPlatform())
		if err != nil {
			return nil, err
		}

		x, err := compute.GetValue(ctx, prodBase0.Image())
		if err != nil {
			return nil, err
		}

		prodBase, err = llbutil.OCILayoutFromImage(ctx, x)
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

	relPath, err := makePkg(bin.GoModulePath, bin.SourcePath)
	if err != nil {
		return nil, err
	}

	goBuild = append(goBuild, relPath)

	state := (llbutil.RunGo{
		Base:       prepareGoMod(base, src, conf.TargetPlatform()).Root(),
		SrcMount:   src,
		WorkingDir: ".",
		Platform:   conf.TargetPlatform(),
	}).With(
		llbutil.PrefixSh(label, conf.TargetPlatform(), "go "+strings.Join(goBuild, " "))...).
		AddMount("/out", prodBase)

	return buildkit.BuildImage(ctx, buildkit.DeferClient(env.Configuration(), conf.TargetPlatform()), conf, state, local)
}

func prepareGoMod(base, src llb.State, platform *specs.Platform) llb.ExecState {
	r := llbutil.RunGo{
		Base:       base,
		SrcMount:   src,
		WorkingDir: ".",
		Platform:   platform,
	}

	ro := llbutil.PrefixSh("updating deps", platform, "go mod download -x")

	if git.AssumeSSHAuth {
		ro = append(ro, llb.AddSSHSocket(llb.SSHID(buildkit.SSHAgentProviderID), llb.SSHSocketTarget("/root/ssh-agent.sock")))
		ro = append(ro, llb.AddEnv("SSH_AUTH_SOCK", "/root/ssh-agent.sock"))
	}

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
		Run(llb.Shlex("apk add --no-cache git openssh"),
			llb.WithCustomName("[prepare build image] apk add --no-cache git openssh")).Root()

	for _, ent := range git.NoPromptEnv() {
		st = st.AddEnv(ent[0], ent[1])
	}

	// Don't block builds on checking the pubkey of the target ssh host.
	// XXX security
	st = st.AddEnv("GIT_SSH_COMMAND", "ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no")

	if llbutil.GitCredentialsBuildkitSecret != "" {
		st = st.Run(llb.Shlex("git config --global credential.helper store")).Root()
	}

	return st
}
