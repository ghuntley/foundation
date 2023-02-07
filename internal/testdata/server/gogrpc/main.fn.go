// This file was automatically generated by Namespace.
// DO NOT EDIT. To update, re-run `ns generate`.
// This file was automatically generated.
package main

import (
	"context"
	"flag"

	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/std/go/server"
)

func main() {
	flag.Parse()

	resources := core.PrepareEnv("namespacelabs.dev/foundation/internal/testdata/server/gogrpc")
	defer resources.Close(context.Background())

	ctx := core.WithResources(context.Background(), resources)

	depgraph := core.NewDependencyGraph()
	RegisterInitializers(depgraph)
	if err := depgraph.RunInitializers(ctx); err != nil {
		core.ZLog.Fatal().Err(err).Send()
	}

	server.InitializationDone()

	server.Listen(ctx, func(srv server.Server) {
		if errs := WireServices(ctx, srv, depgraph); len(errs) > 0 {
			core.ZLog.Fatal().Errs("errors", errs).Msgf("%d services failed to initialize.", len(errs))
		}
	})
}
