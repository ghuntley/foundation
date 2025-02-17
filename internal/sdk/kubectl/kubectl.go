// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package kubectl

import (
	"context"
	"fmt"
	"path/filepath"

	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"namespacelabs.dev/foundation/internal/artifacts"
	"namespacelabs.dev/foundation/internal/artifacts/unpack"
	"namespacelabs.dev/foundation/internal/compute"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/std/tasks"
)

const version = "1.23.6"

var Pins = map[string]artifacts.Reference{
	"linux/amd64": {
		URL: fmt.Sprintf("https://dl.k8s.io/release/v%s/bin/linux/amd64/kubectl", version),
		Digest: schema.Digest{
			Algorithm: "sha256",
			Hex:       "703a06354bab9f45c80102abff89f1a62cbc2c6d80678fd3973a014acc7c500a",
		},
	},
	"linux/arm64": {
		URL: fmt.Sprintf("https://dl.k8s.io/release/v%s/bin/linux/arm64/kubectl", version),
		Digest: schema.Digest{
			Algorithm: "sha256",
			Hex:       "4be771c8e6a082ba61f0367077f480237f9858ef5efe14b1dbbfc05cd42fc360",
		},
	},
	"darwin/arm64": {
		URL: fmt.Sprintf("https://dl.k8s.io/release/v%s/bin/darwin/arm64/kubectl", version),
		Digest: schema.Digest{
			Algorithm: "sha256",
			Hex:       "d03e1f6b88443e46c11f5940a1fa935c91a0d67f5cc4ffeec35083b7e054034d",
		},
	},
	"darwin/amd64": {
		URL: fmt.Sprintf("https://dl.k8s.io/release/v%s/bin/darwin/amd64/kubectl", version),
		Digest: schema.Digest{
			Algorithm: "sha256",
			Hex:       "dedb7784744f581dc7157b0a6589c7e15d4d14a1bcd25dc5876548805034dffe",
		},
	},
}

type Kubectl string

func EnsureSDK(ctx context.Context, p specs.Platform) (Kubectl, error) {
	sdk, err := SDK(ctx, p)
	if err != nil {
		return "", err
	}

	return compute.GetValue(ctx, sdk)
}

func SDK(ctx context.Context, p specs.Platform) (compute.Computable[Kubectl], error) {
	key := fmt.Sprintf("%s/%s", p.OS, p.Architecture)
	ref, ok := Pins[key]
	if !ok {
		return nil, fnerrors.New("platform not supported: %s", key)
	}

	w := unpack.Unpack("kubectl", unpack.MakeFilesystem("kubectl", 0755, ref))

	return compute.Map(
		tasks.Action("kubectl.ensure").Arg("version", version).HumanReadablef("Ensuring kubectl %s is installed", version),
		compute.Inputs().Computable("kubectl", w),
		compute.Output{},
		func(ctx context.Context, r compute.Resolved) (Kubectl, error) {
			return Kubectl(filepath.Join(compute.MustGetDepValue(r, w, "kubectl").Files, "kubectl")), nil
		}), nil
}
