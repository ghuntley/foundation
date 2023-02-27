// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package debug

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/fnapi"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/std/cfg"
)

func NewFnServicesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "fnservices",
		Short:  "Namespace services-related activities (internal only).",
		Hidden: true,
	}

	var fqdn, target string

	mapAddr := fncobra.CmdWithEnv(&cobra.Command{
		Use:   "naming-map",
		Short: "Maps a FQDN within Namespace Cloud's scope to a particular target (e.g. CNAME, or IP address).",
		Args:  cobra.NoArgs,
	}, func(ctx context.Context, env cfg.Context, args []string) error {
		return fnapi.Map(ctx, fqdn, target)
	})

	mapAddr.Flags().StringVar(&fqdn, "fqdn", "", "Fully qualified domain.")
	mapAddr.Flags().StringVar(&target, "target", "", "Target address.")

	_ = mapAddr.MarkFlagRequired("fqdn")
	_ = mapAddr.MarkFlagRequired("target")

	var org string

	allocateName := fncobra.CmdWithEnv(&cobra.Command{
		Use:   "naming-allocate-name",
		Short: "Allocates a TLS certificate within Namespace Cloud's scope.",
		Args:  cobra.NoArgs,
	}, func(ctx context.Context, env cfg.Context, args []string) error {
		if fqdn == "" {
			return fnerrors.BadInputError("--fqdn needs to be specified")
		}

		nr, err := fnapi.AllocateName(ctx, fnapi.AllocateOpts{
			FQDN: fqdn,
			Org:  org,
		})
		if err != nil {
			return err
		}

		w := json.NewEncoder(console.Stdout(ctx))
		w.SetIndent("", "  ")
		return w.Encode(nr)
	})

	allocateName.Flags().StringVar(&fqdn, "fqdn", "", "Fully qualified domain.")
	allocateName.Flags().StringVar(&org, "org", "", "Organization to identify as.")

	cmd.AddCommand(mapAddr)
	cmd.AddCommand(allocateName)

	return cmd
}
