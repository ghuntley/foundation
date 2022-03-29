// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package buf

import (
	"context"

	"namespacelabs.dev/foundation/build/binary"
	"namespacelabs.dev/foundation/internal/artifacts/oci"
	"namespacelabs.dev/foundation/internal/engine/ops"
	"namespacelabs.dev/foundation/runtime/tools"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/workspace"
	"namespacelabs.dev/foundation/workspace/compute"
)

const baseImgPkg schema.PackageName = "namespacelabs.dev/foundation/std/sdk/buf/baseimg"

func init() {
	workspace.StaticDeps = append(workspace.StaticDeps, baseImgPkg)
}

func Image(ctx context.Context, env ops.Environment, loader workspace.Packages) compute.Computable[oci.Image] {
	pkg, err := loader.LoadByName(ctx, baseImgPkg)
	if err != nil {
		return compute.Error[oci.Image](err)
	}

	platform := tools.Impl().HostPlatform()
	prep, err := binary.PlanImage(ctx, pkg, env, true, &platform)
	if err != nil {
		return compute.Error[oci.Image](err)
	}

	return prep.Image
}