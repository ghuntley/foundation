// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
	"namespacelabs.dev/foundation/internal/cli/nsboot"
)

func NewUpdateNSCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Updates ns to the latest version.",
		// There's an implicit contract with DoMain, that it doesn't perform version updates when a `update-ns` alias is present.
		Aliases: []string{"update-ns"},

		RunE: fncobra.RunE(func(ctx context.Context, args []string) error {
			return nsboot.ForceUpdate(ctx, "ns")
		}),
	}

	return cmd
}
