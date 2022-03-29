// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package debug

import (
	"context"
	"strings"

	"github.com/spf13/cobra"
	"namespacelabs.dev/foundation/internal/artifacts"
	"namespacelabs.dev/foundation/internal/artifacts/download"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
	"namespacelabs.dev/foundation/internal/fntypes"
	"namespacelabs.dev/foundation/workspace/compute"
	"namespacelabs.dev/foundation/workspace/tasks"
)

func newDownloadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download",
		Short: "Downloads an URL.",
		Args:  cobra.ArbitraryArgs,

		RunE: fncobra.RunE(func(ctx context.Context, args []string) error {
			var downloads []compute.Computable[compute.ByteStream]

			for _, arg := range args {
				var artifact artifacts.Reference
				parts := strings.SplitN(arg, "@", 2)
				artifact.URL = parts[0]
				if len(parts) == 2 {
					var err error
					artifact.Digest, err = fntypes.ParseDigest(parts[1])
					if err != nil {
						return err
					}
				}

				downloads = append(downloads, download.URL(artifact))
			}

			_, err := compute.Get(ctx, compute.Collect(tasks.Action("download-all"), downloads...))
			if err != nil {
				return err
			}

			return nil
		}),
	}

	return cmd
}