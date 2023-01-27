// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package parsing

import (
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/std/pkggraph"
)

func transformBinary(loc pkggraph.Location, bin *schema.Binary) error {
	if bin.PackageName != "" {
		return fnerrors.NewWithLocation(loc, "package_name can not be set")
	}

	if bin.Name == "" {
		return fnerrors.NewWithLocation(loc, "binary name can't be empty")
	}

	if bin.BuildPlan == nil {
		return fnerrors.NewWithLocation(loc, "a build plan is required")
	}

	bin.PackageName = loc.PackageName.String()

	if len(bin.GetConfig().GetCommand()) == 0 {
		hasGoLayers := false
		for _, layer := range bin.BuildPlan.LayerBuildPlan {
			if isImagePlanGo(layer) {
				hasGoLayers = true
				break
			}
		}

		// For Go, by default, assume the binary is built with the same name as the package name.
		// TODO: revisit this heuristic.
		if hasGoLayers {
			if bin.Config == nil {
				bin.Config = &schema.BinaryConfig{}
			}

			bin.Config.Command = []string{"/" + bin.Name}
		}
	}

	return nil
}

func isImagePlanGo(plan *schema.ImageBuildPlan) bool {
	return plan.GoBuild != nil || plan.GoPackage != ""
}
