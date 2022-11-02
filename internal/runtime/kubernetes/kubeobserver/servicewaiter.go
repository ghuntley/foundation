// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package kubeobserver

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s "k8s.io/client-go/kubernetes"
	"namespacelabs.dev/foundation/framework/kubernetes/kubedef"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/protos"
	"namespacelabs.dev/foundation/internal/runtime/kubernetes/client"
	"namespacelabs.dev/foundation/schema/orchestration"
	"namespacelabs.dev/foundation/schema/runtime"
	"namespacelabs.dev/foundation/std/tasks"
)

const dialTimeout = 100 * time.Millisecond

type serviceWaiter struct {
	namespace, name string

	isReady func(string, []corev1.Pod, error) bool

	mu                    sync.Mutex
	portCount, matchCount int
}

// FormatProgress implements ActionProgress.
func (w *serviceWaiter) FormatProgress() string {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.portCount == 0 {
		return "(waiting for ports...)"
	}

	return fmt.Sprintf("%d / %d", w.matchCount, w.portCount)
}

func (w *serviceWaiter) Prepare(ctx context.Context, c *k8s.Clientset) error {
	tasks.Attachments(ctx).SetProgress(w)
	return nil
}

func (w *serviceWaiter) Poll(ctx context.Context, c *k8s.Clientset) (bool, error) {
	if !client.IsInclusterClient(c) {
		// Emitting this debug message as only incluster deployments know how to determine service readiness.
		fmt.Fprintf(console.Debug(ctx), "will not wait for service %s...\n", w.name)

		// Assume service is always ready for now.
		// TODO implement readiness check that also supports non-incluster deployments.
		return true, nil
	}

	service, err := c.CoreV1().Services(w.namespace).Get(ctx, w.name, v1.GetOptions{})
	if err != nil {
		return false, err
	}

	id, ok := service.Labels[kubedef.K8sServerId]
	if !ok {
		return false, fnerrors.InternalError("service %q is missing server label", w.name)
	}

	scope, ok := service.Annotations[kubedef.K8sServicePackageName]
	if !ok {
		return false, fnerrors.InternalError("service %q is missing package name", w.name)
	}

	pod, err := c.CoreV1().Pods(w.namespace).List(ctx, v1.ListOptions{
		LabelSelector: kubedef.SerializeSelector(map[string]string{
			kubedef.K8sServerId: id,
		}),
	})
	if err != nil {
		return false, err
	}

	var count int
	for _, port := range service.Spec.Ports {
		addr := fmt.Sprintf("%s.%s.svc.cluster.local:%d", service.Name, service.Namespace, port.Port)

		rawConn, err := net.DialTimeout("tcp", addr, dialTimeout)
		if !w.isReady(scope, pod.Items, err) {
			continue
		}

		count++
		rawConn.Close()
	}

	w.mu.Lock()
	defer w.mu.Unlock()
	w.portCount = len(service.Spec.Ports)
	w.matchCount = count

	return w.matchCount > 0 && w.matchCount == w.portCount, nil
}

func WaitForService(namespace, name string, isReady func(string, []corev1.Pod, error) bool) ConditionWaiter[*k8s.Clientset] {
	return &serviceWaiter{namespace: namespace, name: name, isReady: isReady}
}

func WaiterFromServiceErr(err error) *orchestration.Event_WaitStatus {
	return &orchestration.Event_WaitStatus{
		Description: "Service not ready",
		Opaque: protos.WrapAnyOrDie(&runtime.WaitError{
			Message: err.Error(),
		}),
	}
}
