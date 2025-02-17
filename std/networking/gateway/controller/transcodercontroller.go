// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"context"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

const (
	AppliedCondition       = "Applied"
	NormalEvent            = "Normal"
	WarningEvent           = "Warning"
	CreateTranscoder       = "CreateHttpGrpcTranscoder"
	DeleteTranscoder       = "DeleteHttpGrpcTranscoder"
	UpdateTranscoderStatus = "UpdateHttpGrpcTranscoderStatus"
)

type HttpGrpcTranscoderReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	snapshot *TranscoderSnapshot
	recorder record.EventRecorder
}

func (r *HttpGrpcTranscoderReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// Receive the `logr` API from the context.
	log := log.FromContext(ctx)

	transcoder := &HttpGrpcTranscoder{}
	if err := r.Get(ctx, req.NamespacedName, transcoder); err != nil {
		log.Error(err, "Unable to fetch HttpGrpcTranscoder", "namespace", req.Namespace, "name", req.Name)

		if apierrors.IsNotFound(err) {
			r.snapshot.DeleteTranscoder(transcoder)
			// Generate a new envoy snapshot since we have deleted a transcoder.
			if err := r.snapshot.GenerateSnapshot(ctx); err != nil {
				errmsg := "Failed to delete transcoder and generate new envoy snapshot"
				r.recorder.Eventf(transcoder, WarningEvent, DeleteTranscoder,
					"%s for namespace %q and name %q: %v", errmsg, req.Namespace, req.Name, err)
				log.Error(err, errmsg, "namespace", req.Namespace, "name", req.Name)
				return ctrl.Result{}, err
			}
			msg := "Deleted transcoder and successfully generated a new envoy snapshot"
			r.recorder.Eventf(transcoder, NormalEvent, DeleteTranscoder,
				"%s for namespace %q and name %q", msg, req.Namespace, req.Name)
			log.Info(msg, "namespace", req.Namespace, "name", req.Name)
			return ctrl.Result{}, client.IgnoreNotFound(err)
		}
	}

	// Preserve all conditions for the transcoder except "Applied".
	var conditions []metav1.Condition
	for _, c := range transcoder.Status.Conditions {
		if c.Type != AppliedCondition {
			conditions = append(conditions, c)
		}
	}
	appliedCondition := metav1.Condition{
		Type:               AppliedCondition,
		Status:             metav1.ConditionTrue,
		ObservedGeneration: transcoder.GetGeneration(),
		LastTransitionTime: metav1.Now(),
	}

	r.snapshot.AddTranscoder(transcoder)

	// Generate a new envoy snapshot since we have added a transcoder.
	snapshotErr := r.snapshot.GenerateSnapshot(ctx)

	// Update the applied condition if we have an error generating the snapshot.
	if snapshotErr != nil {
		errmsg := "Failed to generate a new envoy snapshot"
		r.recorder.Eventf(transcoder, WarningEvent, CreateTranscoder,
			"%s for namespace %q and name %q: %v", errmsg, req.Namespace, req.Name, snapshotErr)
		log.Error(snapshotErr, errmsg, "namespace", req.Namespace, "name", req.Name)

		appliedCondition.Status = metav1.ConditionFalse
		appliedCondition.Reason = "FailedToGenerateSnapshot"
		appliedCondition.Message = snapshotErr.Error()
	} else {
		msg := "Successfully generated a new envoy snapshot"
		r.recorder.Eventf(transcoder, NormalEvent, CreateTranscoder,
			"%s with version %d for namespace %q and name %q",
			msg, r.snapshot.CurrentSnapshotId(), req.Namespace, req.Name)
		log.Info(msg, "namespace", req.Namespace, "name", req.Name, "version", r.snapshot.CurrentSnapshotId())
	}

	conditions = append(conditions, appliedCondition)
	transcoder.Status.Conditions = conditions

	// Update the status condition on the transcoder.
	if updateErr := r.Client.Status().Update(ctx, transcoder); updateErr != nil {
		// Requeue (rate-limited) if we lost an update race.
		if apierrors.IsConflict(updateErr) {
			msg := "Requeueing since we lost an update race"
			r.recorder.Eventf(transcoder, NormalEvent, UpdateTranscoderStatus,
				"%s for namespace %q and name %q", msg, req.Namespace, req.Name)
			log.Info(msg, "namespace", req.Namespace, "name", req.Name)
			return ctrl.Result{Requeue: true}, nil
		}

		errmsg := "Failed to update transcoder status"
		r.recorder.Eventf(transcoder, WarningEvent, UpdateTranscoderStatus,
			"%s for namespace %q and name %q: %v", errmsg, req.Namespace, req.Name, updateErr)
		log.Error(updateErr, errmsg, "namespace", req.Namespace, "name", req.Name)
		return ctrl.Result{}, updateErr
	} else {
		msg := "Successfully updated transcoder status"
		r.recorder.Eventf(transcoder, NormalEvent, UpdateTranscoderStatus,
			"%s with version %d for namespace %q and name %q",
			msg, r.snapshot.CurrentSnapshotId(), req.Namespace, req.Name)
		log.Info(msg, "namespace", req.Namespace, "name", req.Name, "version", r.snapshot.CurrentSnapshotId())

	}

	return ctrl.Result{}, snapshotErr
}

func (r *HttpGrpcTranscoderReconciler) SetupWithManager(mgr ctrl.Manager, matchNamespace string) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&HttpGrpcTranscoder{}, builder.WithPredicates(predicate.NewPredicateFuncs(
			func(object client.Object) bool {
				return object.GetNamespace() == matchNamespace
			},
		))).
		Complete(r)
}
