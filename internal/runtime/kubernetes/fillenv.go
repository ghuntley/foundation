// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package kubernetes

import (
	"context"
	"strings"

	applycorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	"namespacelabs.dev/foundation/framework/kubernetes/kubedef"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/runtime"
	"namespacelabs.dev/foundation/schema"
	rtschema "namespacelabs.dev/foundation/schema/runtime"
	runtimepb "namespacelabs.dev/foundation/schema/runtime"
)

func fillEnv(ctx context.Context, rt *runtimepb.RuntimeConfig, container *applycorev1.ContainerApplyConfiguration, env []*schema.BinaryConfig_EnvEntry, secrets *secretCollector, ensure *kubedef.EnsureDeployment) (*applycorev1.ContainerApplyConfiguration, error) {
	for _, kv := range env {
		var entry *applycorev1.EnvVarApplyConfiguration

		switch {
		case kv.ExperimentalFromSecret != "":
			parts := strings.SplitN(kv.ExperimentalFromSecret, ":", 2)
			if len(parts) < 2 {
				return nil, fnerrors.New("invalid experimental_from_secret format")
			}
			entry = applycorev1.EnvVar().WithName(kv.Name).
				WithValueFrom(applycorev1.EnvVarSource().WithSecretKeyRef(
					applycorev1.SecretKeySelector().WithName(parts[0]).WithKey(parts[1])))

		case kv.ExperimentalFromDownwardsFieldPath != "":
			entry = applycorev1.EnvVar().WithName(kv.Name).
				WithValueFrom(applycorev1.EnvVarSource().WithFieldRef(
					applycorev1.ObjectFieldSelector().WithFieldPath(kv.ExperimentalFromDownwardsFieldPath)))

		case kv.FromSecretRef != nil:
			if secrets == nil {
				return nil, fnerrors.InternalError("can't use FromSecretRef in this context")
			}

			alloc, err := secrets.allocate(ctx, kv.FromSecretRef)
			if err != nil {
				return nil, err
			}

			entry = applycorev1.EnvVar().WithName(kv.Name).
				WithValueFrom(applycorev1.EnvVarSource().WithSecretKeyRef(
					applycorev1.SecretKeySelector().WithName(alloc.Name).WithKey(alloc.Key),
				))

		case kv.FromServiceEndpoint != nil:
			endpoint, err := runtime.SelectServiceValue(rt, kv.FromServiceEndpoint, runtime.SelectServiceEndpoint)
			if err != nil {
				return nil, err
			}

			entry = applycorev1.EnvVar().WithName(kv.Name).WithValue(endpoint)

		case kv.FromServiceIngress != nil:
			url, err := runtime.SelectServiceValue(rt, kv.FromServiceIngress, runtime.SelectServiceIngress)
			if err != nil {
				return nil, err
			}

			entry = applycorev1.EnvVar().WithName(kv.Name).WithValue(url)

		case kv.FromResourceField != nil:
			ensure.SetContainerFields = append(ensure.SetContainerFields, &rtschema.SetContainerField{
				SetEnv: []*rtschema.SetContainerField_SetValue{
					{
						ContainerName:               *container.Name,
						Key:                         kv.Name,
						Value:                       rtschema.SetContainerField_RESOURCE_CONFIG_FIELD_SELECTOR,
						ResourceConfigFieldSelector: kv.FromResourceField,
					},
				},
			})

			// No environment variable is injected here yet, it will be then patched in by OpEnsureDeployment.

		case kv.FromFieldSelector != nil:
			return nil, fnerrors.BadInputError("fromFieldSelector is not supported in this context")

		default:
			entry = applycorev1.EnvVar().WithName(kv.Name).WithValue(kv.Value)
		}

		if entry != nil {
			container = container.WithEnv(entry)
		}
	}

	return container, nil
}
