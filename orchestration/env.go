// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package orchestration

import (
	"context"

	"google.golang.org/protobuf/proto"
	"namespacelabs.dev/foundation/build/binary"
	"namespacelabs.dev/foundation/internal/fnapi"
	"namespacelabs.dev/foundation/internal/protos"
	"namespacelabs.dev/foundation/runtime/kubernetes/client"
	"namespacelabs.dev/foundation/runtime/kubernetes/kubedef"
	"namespacelabs.dev/foundation/runtime/kubernetes/kubetool"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/std/planning"
)

const (
	serverPkg = "namespacelabs.dev/foundation/orchestration/server"
	toolPkg   = "namespacelabs.dev/foundation/orchestration/server/tool"
)

func MakeSyntheticConfiguration(wsproto *schema.Workspace, envName string, hostEnv *client.HostEnv, extra ...proto.Message) planning.Configuration {
	messages := []proto.Message{hostEnv}
	messages = append(messages, extra...)

	ws := planning.MakeWorkspace(wsproto, nil)

	return planning.MakeConfigurationWith(envName, ws, planning.ConfigurationSlice{Configuration: protos.WrapAnysOrDie(messages...)})
}

func MakeSyntheticContext(wsproto *schema.Workspace, env *schema.Environment, hostEnv *client.HostEnv, extra ...proto.Message) planning.Context {
	cfg := MakeSyntheticConfiguration(wsproto, env.Name, hostEnv, extra...)
	return planning.MakeUnverifiedContext(cfg, env)
}

func MakeOrchestratorContext(ctx context.Context, conf planning.Configuration) (planning.Context, error) {
	// We use a static environment here, since the orchestrator has global scope.
	envProto := &schema.Environment{
		Name:      kubedef.AdminNamespace,
		Runtime:   "kubernetes", // XXX should be an input.
		Ephemeral: false,

		// TODO - this can't be empty, since std/runtime/kubernetes/extension.cue checks it.
		Purpose: schema.Environment_PRODUCTION,
	}

	var prebuilts []*schema.Workspace_BinaryDigest

	for _, pkg := range []string{serverPkg, toolPkg} {
		res, err := fnapi.GetLatestPrebuilt(ctx, schema.PackageName(pkg))
		if err != nil {
			return nil, err
		}

		prebuilts = append(prebuilts, &schema.Workspace_BinaryDigest{
			PackageName: pkg,
			Repository:  res.Repository,
			Digest:      res.Digest,
		})
	}

	cfg := conf.Derive(kubedef.AdminNamespace, func(previous planning.ConfigurationSlice) planning.ConfigurationSlice {
		previous.Configuration = append(previous.Configuration, protos.WrapAnysOrDie(
			&kubetool.KubernetesEnv{Namespace: kubedef.AdminNamespace}, // pin deployments to admin namespace
			&binary.Prebuilts{PrebuiltBinary: prebuilts},
		)...)
		return previous
	})

	return planning.MakeUnverifiedContext(cfg, envProto), nil
}
