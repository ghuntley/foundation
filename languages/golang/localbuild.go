// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package golang

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"namespacelabs.dev/foundation/build"
	"namespacelabs.dev/foundation/internal/artifacts/oci"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/engine/ops"
	"namespacelabs.dev/foundation/internal/fnfs"
	"namespacelabs.dev/foundation/internal/localexec"
	"namespacelabs.dev/foundation/internal/production"
	"namespacelabs.dev/foundation/internal/sdk/golang"
	"namespacelabs.dev/foundation/internal/wscontents"
	"namespacelabs.dev/foundation/workspace/compute"
	"namespacelabs.dev/foundation/workspace/devhost"
	"namespacelabs.dev/foundation/workspace/dirs"
	"namespacelabs.dev/foundation/workspace/tasks"
)

func Build(ctx context.Context, env ops.Environment, bin GoBinary, conf build.Configuration) (compute.Computable[oci.Image], error) {
	return buildLocalImage(ctx, env, conf.Workspace, bin, *conf.Target)
}

func buildLocalImage(ctx context.Context, env ops.Environment, workspace build.Workspace, bin GoBinary, target specs.Platform) (compute.Computable[oci.Image], error) {
	sdk, err := golang.SDK(bin.GoVersion, golang.HostPlatform())
	if err != nil {
		return nil, err
	}

	baseImage, err := production.ServerImage(production.Distroless, target)
	if err != nil {
		return nil, err
	}

	layers := []compute.Computable[oci.Layer]{
		// By depending on workspace.Contents we both get continued updates on changes to the workspace,
		// but also are guaranteed to only be invoked after generation functions run.
		oci.MakeLayer("binary", &compilation{
			sdk:       sdk,
			workspace: workspace.VersionedFS(bin.GoModulePath, bin.isFocus),
			binary:    bin,
			platform:  target,
		}),
	}

	return compute.Named(tasks.Action("go.make-binary-image").Arg("binary", bin), oci.MakeImage(baseImage, layers...)), nil
}

func platformToEnv(platform specs.Platform, cgo int) []string {
	return []string{fmt.Sprintf("CGO_ENABLED=%d", cgo), "GOOS=" + platform.OS, "GOARCH=" + platform.Architecture}
}

func compile(ctx context.Context, sdk golang.LocalSDK, absWorkspace string, targetDir string, bin GoBinary, platform specs.Platform) error {
	env := platformToEnv(platform, 0)
	env = append(env, "GOROOT="+sdk.GoRoot())
	env = append(env, goPrivate())

	if platform.Architecture == "arm" {
		v, err := goarm(platform)
		if err != nil {
			return err
		}
		env = append(env, v)
	}

	modulePath := filepath.Join(absWorkspace, bin.GoModulePath)
	out := filepath.Join(targetDir, bin.BinaryName)
	pkg := makePkg(bin.SourcePath)

	var cmd localexec.Command
	cmd.Label = "go build"
	cmd.Command = sdk.GoBin()
	cmd.Args = []string{"build", "-v", "-o=" + out, pkg}
	cmd.AdditionalEnv = env
	cmd.Dir = modulePath
	return cmd.Run(ctx)
}

func makePkg(srcPath string) string {
	if srcPath == "" || srcPath == "." {
		return "./"
	}

	return "./" + srcPath
}

func goarm(platform specs.Platform) (string, error) {
	if platform.Architecture != "arm" {
		return "", fmt.Errorf("not arm: %v", platform.Architecture)
	}
	v := platform.Variant
	if len(v) != 2 {
		return "", fmt.Errorf("unexpected varient: %v", v)
	}
	if v[0] != 'v' || !('0' <= v[1] && v[1] <= '9') {
		return "", fmt.Errorf("unexpected varient: %v", v)
	}
	return string(v[1]), nil
}

type compilation struct {
	sdk       compute.Computable[golang.LocalSDK]
	workspace compute.Computable[wscontents.Versioned] // We depend on `workspace` so we trigger a re-build on workspace changes.
	binary    GoBinary
	platform  specs.Platform

	compute.LocalScoped[fs.FS]
}

func (c *compilation) Action() *tasks.ActionEvent {
	return tasks.Action("go.build.binary").
		WellKnown(tasks.WkModule, c.binary.ModuleName).
		Arg("binary", c.binary.BinaryName).
		Arg("module_path", c.binary.GoModulePath).
		Arg("source_path", c.binary.SourcePath).
		Arg("platform", devhost.FormatPlatform(c.platform))
}

func (c *compilation) Inputs() *compute.In {
	in := compute.Inputs().
		JSON("binary", c.binary).
		JSON("platform", c.platform).
		Computable("workspace", c.workspace).
		Computable("sdk", c.sdk)
	if !c.binary.UnsafeCacheable {
		in = in.Indigestible("localfs", nil)
	}
	return in
}

func (c *compilation) Compute(ctx context.Context, deps compute.Resolved) (fs.FS, error) {
	w := compute.GetDepValue(deps, c.workspace, "workspace")
	sdk := compute.GetDepValue(deps, c.sdk, "sdk")

	targetDir, err := dirs.CreateUserTempDir("go", "build")
	if err != nil {
		return nil, err
	}

	if err := compile(ctx, sdk, w.Abs(), targetDir, c.binary, c.platform); err != nil {
		return nil, err
	}

	result := fnfs.Local(targetDir)

	// Only initiate a cleanup after we're done compiling.
	compute.On(ctx).Cleanup(tasks.Action("go.build.cleanup"), func(ctx context.Context) error {
		if err := os.RemoveAll(targetDir); err != nil {
			fmt.Fprintln(console.Warnings(ctx), "failed to cleanup target dir", err)
		}
		return nil // Never fail.
	})

	return result, nil
}