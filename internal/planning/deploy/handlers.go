// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package deploy

import (
	"context"
	"sort"
	"strings"

	"namespacelabs.dev/foundation/internal/planning"
	"namespacelabs.dev/foundation/internal/planning/invocation"
	"namespacelabs.dev/foundation/internal/planning/tool"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/std/pkggraph"
)

type handlerSource struct {
	Server  planning.Server
	Package schema.PackageName
	Plan    pkggraph.PreparedProvisionPlan
}

func computeHandlers(ctx context.Context, in *planning.Stack) ([]*tool.Definition, error) {
	var handlers []*tool.Definition

	var sources []handlerSource
	for _, ps := range in.Servers {
		srv := ps.Server
		for _, n := range ps.ParsedDeps {
			sources = append(sources, handlerSource{srv, n.Package.PackageName(), n.ProvisionPlan.PreparedProvisionPlan})
		}
		sources = append(sources, handlerSource{srv, srv.PackageName(), srv.Provisioning})
	}

	for _, src := range sources {
		h, err := parseHandlers(ctx, src.Server, src.Plan)
		if err != nil {
			return nil, err
		}
		handlers = append(handlers, h...)
	}

	sort.SliceStable(handlers, func(i, j int) bool {
		if handlers[i].TargetServer == handlers[j].TargetServer {
			return strings.Compare(handlers[i].Source.PackageName.String(), handlers[j].Source.PackageName.String()) < 0
		}

		return strings.Compare(handlers[i].TargetServer.String(), handlers[j].TargetServer.String()) < 0
	})

	return handlers, nil
}

func parseHandlers(ctx context.Context, server planning.Server, pr pkggraph.PreparedProvisionPlan) ([]*tool.Definition, error) {
	source := tool.Source{
		PackageName: server.PackageName(),
	}

	var handlers []*tool.Definition
	for _, dec := range pr.ComputePlanWith {
		invocation, err := invocation.BuildAndPrepare(ctx, server.SealedContext(), server.SealedContext(), &server.Location, dec)
		if err != nil {
			return nil, err
		}

		handlers = append(handlers, &tool.Definition{
			TargetServer: server.PackageName(),
			Source:       source,
			Invocation:   invocation,
		})
	}

	return handlers, nil
}
