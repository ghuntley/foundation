// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package fnapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/atomic"
	"gotest.tools/assert"
)

func TestTelemetryDisabled(t *testing.T) {
	reset := setupEnv(t)
	defer reset()

	tel := InternalNewTelemetry(context.Background(), generateTestIDs)
	tel.enabled = false
	tel.errorLogging = true

	cmd := &cobra.Command{
		Use: "fake-command",
		Run: func(cmd *cobra.Command, args []string) {
			tel.RecordInvocation(context.Background(), cmd, args)
		}}

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("Calls to TelemetryService are fobidden when telemetry is disabled.")
	}))
	defer svr.Close()

	tel.backendAddress = svr.URL

	_ = cmd.Execute()
	tel.RecordError(context.Background(), fmt.Errorf("foo error"))

	// Due to the async http server nature it may not have time to handle the request.
	time.Sleep(time.Millisecond * 100)
}

func generateTestIDs(ctx context.Context) (ephemeralCliID, bool) {
	return ephemeralCliID{newRandID(), newRandID()}, true
}

func TestTelemetryDisabledViaEnv(t *testing.T) {
	reset := setupEnv(t)
	defer reset()

	tel := InternalNewTelemetry(context.Background(), generateTestIDs)
	tel.enabled = true
	tel.errorLogging = true

	t.Setenv("DO_NOT_TRACK", "1")

	cmd := &cobra.Command{
		Use: "fake-command",
		Run: func(cmd *cobra.Command, args []string) {
			tel.RecordInvocation(context.Background(), cmd, args)
		}}

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.EscapedPath(), "/telemetry.TelemetryService") {
			t.Errorf("Calls to TelemetryService are fobidden when telemetry is disabled.")
		}

	}))
	defer svr.Close()

	tel.backendAddress = svr.URL

	_ = cmd.Execute()
	tel.RecordError(context.Background(), fmt.Errorf("foo error"))

	// Due to the async http server nature it may not have time to handle the request.
	time.Sleep(time.Millisecond * 100)
}

func TestTelemetryDisabledViaViper(t *testing.T) {
	reset := setupEnv(t)
	defer reset()

	viper.Set("telemetry", false)

	tel := InternalNewTelemetry(context.Background(), generateTestIDs)
	tel.enabled = true
	tel.errorLogging = true

	cmd := &cobra.Command{
		Use: "fake-command",
		Run: func(cmd *cobra.Command, args []string) {
			tel.RecordInvocation(context.Background(), cmd, args)
		}}

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.EscapedPath(), "/telemetry.TelemetryService") {
			t.Errorf("Calls to TelemetryService are fobidden when telemetry is disabled.")
		}

	}))
	defer svr.Close()

	tel.backendAddress = svr.URL

	_ = cmd.Execute()
	tel.RecordError(context.Background(), fmt.Errorf("foo error"))

	// Due to the async http server nature it may not have time to handle the request.
	time.Sleep(time.Millisecond * 100)
}

func TestTelemetryRecordInvocationAnon(t *testing.T) {
	reset := setupEnv(t)
	defer reset()

	tel := InternalNewTelemetry(context.Background(), generateTestIDs)
	tel.enabled = true
	tel.errorLogging = true

	sentID := make(chan string, 1)
	cmd := &cobra.Command{
		Use: "fake-command",
		Run: func(cmd *cobra.Command, args []string) {
			defer close(sentID)
			sentID <- tel.RecordInvocation(context.Background(), cmd, args)
		}}
	cmd.PersistentFlags().Bool("dummy_flag", false, "")

	fakeArg := "fake/arg/value"
	fakeArgs := []string{"--dummy_flag", fakeArg}
	cmd.SetArgs(fakeArgs)

	var req recordInvocationRequest

	receivedID := make(chan string, 1)
	svr := httptest.NewServer(assertGrpcInvocation(t, "/telemetry.TelemetryService/RecordInvocation", &req, func(w http.ResponseWriter) {
		defer close(receivedID)

		assert.Equal(t, req.Command, cmd.Use, req)

		// Assert that we don't transmit user data in plain text.
		assert.Equal(t, len(req.Arg), 1, req)
		assert.Assert(t, req.Arg[0].Hash != fakeArg, req)
		assert.Equal(t, req.Arg[0].Plaintext, "", req)
		assert.Equal(t, len(req.Flag), 1, req)
		assert.Equal(t, req.Flag[0].Name, "dummy_flag", req)
		assert.Assert(t, req.Flag[0].Hash != "true", req)
		assert.Equal(t, req.Flag[0].Plaintext, "", req)

		receivedID <- req.ID
	}))

	defer svr.Close()

	tel.backendAddress = svr.URL

	err := cmd.Execute()
	assert.NilError(t, err)

	assert.Equal(t, <-receivedID, <-sentID) // Make sure we validated the request.
}

func TestTelemetryRecordErrorPlaintext(t *testing.T) {
	reset := setupEnv(t)
	defer reset()

	tel := InternalNewTelemetry(context.Background(), generateTestIDs)
	tel.enabled = true
	tel.errorLogging = true
	tel.recID = atomic.NewString("fake-id")

	var req recordErrorRequest
	receivedID := make(chan string, 1)
	svr := httptest.NewServer(assertGrpcInvocation(t, "/telemetry.TelemetryService/RecordError", &req, func(_ http.ResponseWriter) {
		defer close(receivedID)

		assert.Assert(t, req.Message != "", req)

		receivedID <- req.ID
	}))
	defer svr.Close()

	tel.backendAddress = svr.URL

	tel.RecordError(context.Background(), fmt.Errorf("foo error"))

	// Assert on intercepted request outside the HandlerFunc to ensure the handler is called
	assert.Equal(t, <-receivedID, tel.recID.Load())
}

func assertGrpcInvocation(t *testing.T, method string, request interface{}, handle func(http.ResponseWriter)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		if r.Method != "POST" {
			t.Errorf("expected method=POST, got method=%v", r.Method)
		}

		bodyBytes, err := io.ReadAll(r.Body)
		assert.NilError(t, err)

		if r.URL.EscapedPath() == method {
			err := json.Unmarshal(bodyBytes, request)
			assert.NilError(t, err)
			handle(rw)
		} else {
			t.Errorf("expected url=%q, got url=%q", method, r.URL.EscapedPath())
		}
	}
}

func setupEnv(t *testing.T) func() {
	t.Setenv("DO_NOT_TRACK", "")
	t.Setenv("CI", "")

	viper.Set("telemetry", true)

	return func() { viper.Reset() }
}
