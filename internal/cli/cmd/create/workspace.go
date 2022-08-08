// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package create

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/console/tui"
	"namespacelabs.dev/foundation/internal/fnfs"
	"namespacelabs.dev/foundation/internal/frontend/cuefrontend"
	"namespacelabs.dev/foundation/internal/git"
)

const (
	workspaceFileTemplate = `module: "%s"
`
	vscodeExtensionsFilePath = ".vscode/extensions.json"
	vscodeExtensionsTemplate = `{
    "recommendations": [
        "golang.go",
        "esbenp.prettier-vscode",
        "zxh404.vscode-proto3",
        "namespacelabs.namespace-vscode"
    ]
}`
	gitignoreFilePath = ".gitignore"
	gitignoreTemplate = `# Namespace configuration of this specific host.
devhost.textpb

# Typescript/Node.js/Yarn
node_modules
**/.yarn/*
!**/.yarn/patches
!**/.yarn/plugins
!**/.yarn/releases
!**/.yarn/sdks
!**/.yarn/versions
`
	gitpodFilePath = ".gitpod.yml"

	// `ns create starter --dry_run` prints to the terminal in VSCode what we have done for the user.
	gitpodTemplate = `image: us-docker.pkg.dev/foundation-344819/prebuilts/namespacelabs.dev/foundation/internal/gitpod/pinned@sha256:b7dc2bbda2aa97c17f875a55c95b38dafad8a7ec933e34fdd88a0185caeac18e
tasks:
  - name: prepare
    init: |
      ns login
      ns prepare new-cluster
      ns create starter --dry_run --suggest_prepare=false

ports:
- name: Namespace Dev UI
  port: 4001
  onOpen: open-preview
`
)

func newWorkspaceCmd(runCommand func(ctx context.Context, args []string) error) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "workspace [directory]",
		Short: "Initializes a workspace.",
	}

	cmd.RunE = fncobra.RunE(func(ctx context.Context, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		fsfs := fnfs.ReadWriteLocalFS(cwd)

		if err := writeWorkspaceConfig(ctx, fsfs, args); err != nil {
			return err
		}
		if err := writeFileIfDoesntExist(ctx, console.Stdout(ctx), fsfs, vscodeExtensionsFilePath, vscodeExtensionsTemplate); err != nil {
			return err
		}
		if err := updateGitignore(ctx, console.Stdout(ctx), fsfs); err != nil {
			return err
		}

		if isRoot, err := git.IsRepoRoot(ctx); err == nil && isRoot {
			if err := writeFileIfDoesntExist(ctx, console.Stdout(ctx), fsfs, gitpodFilePath, gitpodTemplate); err != nil {
				return err
			}
		}

		return runCommand(ctx, []string{"tidy"})
	})

	return cmd
}

func askWorkspaceName(ctx context.Context) (string, error) {
	placeholder := "github.com/username/reponame"
	if url, err := git.RemoteUrl(ctx); err == nil {
		placeholder = url
	}

	return tui.Ask(ctx,
		"Workspace name?",
		"The workspace name should to match the Github repository name.",
		placeholder)
}

func workspaceNameFromArgs(ctx context.Context, args []string) (string, error) {
	if len(args) == 0 {
		workspaceName, err := askWorkspaceName(ctx)
		if err != nil {
			return "", err
		}
		if workspaceName == "" {
			return "", context.Canceled
		}
		return workspaceName, nil
	} else {
		return args[0], nil
	}
}

func writeWorkspaceConfig(ctx context.Context, fsfs fnfs.ReadWriteFS, args []string) error {
	workspaceName, err := workspaceNameFromArgs(ctx, args)
	if err != nil {
		return err
	}
	if workspaceName == "" {
		return context.Canceled
	}

	return writeFileIfDoesntExist(ctx, nil, fsfs, cuefrontend.WorkspaceFile, fmt.Sprintf(workspaceFileTemplate, workspaceName))
}

func writeFileIfDoesntExist(ctx context.Context, out io.Writer, fsfs fnfs.ReadWriteFS, fn string, content string) error {
	if f, err := fsfs.Open(fn); err == nil {
		f.Close()
		return nil
	}

	return fnfs.WriteWorkspaceFile(ctx, out, fsfs, fn, func(w io.Writer) error {
		_, err := fmt.Fprint(w, content)
		return err
	})
}

func updateGitignore(ctx context.Context, out io.Writer, fsfs fnfs.ReadWriteFS) error {
	f, err := fsfs.Open(gitignoreFilePath)
	if err != nil {
		// file does not exist

		return fnfs.WriteWorkspaceFile(ctx, out, fsfs, gitignoreFilePath, func(w io.Writer) error {
			_, err := fmt.Fprint(w, gitignoreTemplate)
			return err
		})
	}

	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	content := string(data)

	if strings.Contains(content, "devhost.textpb") {
		// Found Namespace bits in .gitignore. Don't touch it.
		return nil
	}
	if strings.Contains(content, "yarn") {
		// .gitignore is apparently not generated by us but already has some yarn bits.
		// Let's not touch it but just say what we'd recommend.
		fmt.Fprintf(console.Stdout(ctx), "Found existing .gitignore file. Consider adding:\n# Namespace .gitignore begin\n%s\n# Namespace .gitignore end\n", gitignoreTemplate)
		return nil
	}

	content = fmt.Sprintf("%s\n%s", content, gitignoreTemplate)

	return fnfs.WriteWorkspaceFile(ctx, out, fsfs, gitignoreFilePath, func(w io.Writer) error {
		_, err := fmt.Fprint(w, content)
		return err
	})
}
