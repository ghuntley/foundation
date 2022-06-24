// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package main

import (
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
	"namespacelabs.dev/foundation/cmd/nspipelines/cmd/github"
	"namespacelabs.dev/foundation/cmd/nspipelines/cmd/workspace"
	"namespacelabs.dev/foundation/workspace/tasks"
	"namespacelabs.dev/foundation/workspace/tasks/simplelog"
)

const maxLogLevel = 0

func main() {
	root := &cobra.Command{
		Use: "nspipelines",

		TraverseChildren: true,
	}

	root.AddCommand(github.NewGithubCmd())
	root.AddCommand(workspace.NewWorkspaceCmd())

	ctx := tasks.WithSink(context.Background(), simplelog.NewSink(os.Stderr, maxLogLevel))

	if err := root.ExecuteContext(ctx); err != nil {
		log.Fatal(err)
	}
}
