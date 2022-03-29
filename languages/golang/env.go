// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package golang

import (
	"fmt"
	"os"
)

func goPrivate() string {
	// All namespace repositories are private at the moment. Go needs to know this
	// in order to use git, rather than the http go module proxy, to fetch
	// dependencies. Because the user may have a GOPRIVATE configuration themselves,
	// we append to an existing configuration if one exists.
	const namespaceRepos = "namespacelabs.dev/*"
	if existing := os.Getenv("GOPRIVATE"); existing != "" {
		return fmt.Sprintf("GOPRIVATE=%s,%s", existing, namespaceRepos)
	}
	return "GOPRIVATE=" + namespaceRepos
}