// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package kubeobserver

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s "k8s.io/client-go/kubernetes"
	"namespacelabs.dev/foundation/internal/fnerrors"
)

func determineDeploymentPodSelector(ctx context.Context, cli *k8s.Clientset, ns string, owner string, gen int64) (map[string]string, error) {
	// TODO explore how to limit the list here (e.g. through labels or by using a different API)
	replicasets, err := cli.AppsV1().ReplicaSets(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fnerrors.InvocationError("kubernetes", "unable to list replica sets: %w", err)
	}

	for _, replicaset := range replicasets.Items {
		if replicaset.ObjectMeta.Annotations["deployment.kubernetes.io/revision"] != fmt.Sprintf("%d", gen) {
			continue
		}

		for _, o := range replicaset.ObjectMeta.OwnerReferences {
			if o.Name == owner {
				if hash, ok := replicaset.Labels["pod-template-hash"]; ok {
					return map[string]string{
						"pod-template-hash": hash,
					}, nil
				}

				return nil, nil
			}
		}
	}

	return nil, nil
}
