// This file was automatically generated by Namespace.
// DO NOT EDIT. To update, re-run `ns generate`.

package multidb

import (
	"context"
	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/std/go/server"
	"namespacelabs.dev/foundation/universe/db/postgres"
	"namespacelabs.dev/foundation/universe/db/postgres/incluster"
	"namespacelabs.dev/foundation/universe/db/postgres/rds"
)

// Dependencies that are instantiated once for the lifetime of the service.
type ServiceDeps struct {
	Postgres *postgres.DB
	Rds      *postgres.DB
}

// Verify that WireService is present and has the appropriate type.
type checkWireService func(context.Context, server.Registrar, ServiceDeps)

var _ checkWireService = WireService

var (
	Package__2q8a4u = &core.Package{
		PackageName: "namespacelabs.dev/foundation/internal/testdata/service/multidb",
	}

	Provider__2q8a4u = core.Provider{
		Package:     Package__2q8a4u,
		Instantiate: makeDeps__2q8a4u,
	}
)

func makeDeps__2q8a4u(ctx context.Context, di core.Dependencies) (_ interface{}, err error) {
	var deps ServiceDeps

	if err := di.Instantiate(ctx, incluster.Provider__udoubi, func(ctx context.Context, v interface{}) (err error) {
		// name: "postgreslist"
		if deps.Postgres, err = incluster.ProvideDatabase(ctx, core.MustUnwrapProto("Cgxwb3N0Z3Jlc2xpc3Q=", &incluster.Database{}).(*incluster.Database), v.(incluster.ExtensionDeps)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	if err := di.Instantiate(ctx, rds.Provider__4j13h1, func(ctx context.Context, v interface{}) (err error) {
		// name: "postgreslist"
		if deps.Rds, err = rds.ProvideDatabase(ctx, core.MustUnwrapProto("Cgxwb3N0Z3Jlc2xpc3Q=", &rds.Database{}).(*rds.Database), v.(rds.ExtensionDeps)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return deps, nil
}
