// This file was automatically generated.
package main

import (
	"context"

	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/std/go/grpc/interceptors"
	"namespacelabs.dev/foundation/std/go/grpc/metrics"
	"namespacelabs.dev/foundation/std/go/grpc/server"
	"namespacelabs.dev/foundation/std/monitoring/tracing"
	"namespacelabs.dev/foundation/std/secrets"
	"namespacelabs.dev/foundation/std/testdata/service/list"
	"namespacelabs.dev/foundation/universe/db/postgres/creds"
	"namespacelabs.dev/foundation/universe/db/postgres/incluster"
)

type ServerDeps struct {
	list list.ServiceDeps
}

func PrepareDeps(ctx context.Context) (*ServerDeps, error) {
	var server ServerDeps
	var di core.DepInitializer
	var metrics0 metrics.ExtensionDeps

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/go/grpc/interceptors",
		Instance:    "metrics0",
		Do: func(ctx context.Context) (err error) {
			if metrics0.Interceptors, err = interceptors.ProvideInterceptorRegistration(ctx, "namespacelabs.dev/foundation/std/go/grpc/metrics", nil); err != nil {
				return err
			}
			return nil
		},
	})

	var tracing0 tracing.ExtensionDeps

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/go/grpc/interceptors",
		Instance:    "tracing0",
		Do: func(ctx context.Context) (err error) {
			if tracing0.Interceptors, err = interceptors.ProvideInterceptorRegistration(ctx, "namespacelabs.dev/foundation/std/monitoring/tracing", nil); err != nil {
				return err
			}
			return nil
		},
	})

	var creds0 creds.ExtensionDeps

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/secrets",
		Instance:    "creds0",
		Do: func(ctx context.Context) (err error) {
			// name: "postgres_password_file"
			// provision: PROVISION_INLINE
			// provision: PROVISION_AS_FILE
			p := &secrets.Secret{}
			core.MustUnwrapProto("ChZwb3N0Z3Jlc19wYXNzd29yZF9maWxlEgIBAg==", p)

			if creds0.Password, err = secrets.ProvideSecret(ctx, "namespacelabs.dev/foundation/universe/db/postgres/creds", p); err != nil {
				return err
			}
			return nil
		},
	})

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/secrets",
		Instance:    "creds0",
		Do: func(ctx context.Context) (err error) {
			// name: "postgres_user_file"
			// provision: PROVISION_INLINE
			// provision: PROVISION_AS_FILE
			p := &secrets.Secret{}
			core.MustUnwrapProto("ChJwb3N0Z3Jlc191c2VyX2ZpbGUSAgEC", p)

			if creds0.User, err = secrets.ProvideSecret(ctx, "namespacelabs.dev/foundation/universe/db/postgres/creds", p); err != nil {
				return err
			}
			return nil
		},
	})

	var incluster0 incluster.ExtensionDeps

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/universe/db/postgres/creds",
		Instance:    "incluster0",
		DependsOn:   []string{"creds0"}, Do: func(ctx context.Context) (err error) {
			if incluster0.Creds, err = creds.ProvideCreds(ctx, "namespacelabs.dev/foundation/universe/db/postgres/incluster", nil, creds0); err != nil {
				return err
			}
			return nil
		},
	})

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/go/core",
		Instance:    "incluster0",
		Do: func(ctx context.Context) (err error) {
			if incluster0.ReadinessCheck, err = core.ProvideReadinessCheck(ctx, "namespacelabs.dev/foundation/universe/db/postgres/incluster", nil); err != nil {
				return err
			}
			return nil
		},
	})

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/universe/db/postgres/incluster",
		Instance:    "server.list",
		DependsOn:   []string{"incluster0"}, Do: func(ctx context.Context) (err error) {
			// name: "list"
			// schema_file: {
			//   path: "schema.sql"
			//   contents: "CREATE TABLE IF NOT EXISTS list (\n    Id INT GENERATED ALWAYS AS IDENTITY,\n    Item varchar(255) NOT NULL,\n    PRIMARY KEY(Id)\n);"
			// }
			p := &incluster.Database{}
			core.MustUnwrapProto("CgRsaXN0EpABCgpzY2hlbWEuc3FsEoEBQ1JFQVRFIFRBQkxFIElGIE5PVCBFWElTVFMgbGlzdCAoCiAgICBJZCBJTlQgR0VORVJBVEVEIEFMV0FZUyBBUyBJREVOVElUWSwKICAgIEl0ZW0gdmFyY2hhcigyNTUpIE5PVCBOVUxMLAogICAgUFJJTUFSWSBLRVkoSWQpCik7", p)

			if server.list.Db, err = incluster.ProvideDatabase(ctx, "namespacelabs.dev/foundation/std/testdata/service/list", p, incluster0); err != nil {
				return err
			}
			return nil
		},
	})

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/go/grpc/metrics",
		DependsOn:   []string{"metrics0"},
		Do: func(ctx context.Context) error {
			return metrics.Prepare(ctx, metrics0)
		},
	})

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/monitoring/tracing",
		DependsOn:   []string{"tracing0"},
		Do: func(ctx context.Context) error {
			return tracing.Prepare(ctx, tracing0)
		},
	})

	return &server, di.Wait(ctx)
}

func WireServices(ctx context.Context, srv *server.Grpc, server *ServerDeps) {
	list.WireService(ctx, srv, server.list)
	srv.RegisterGrpcGateway(list.RegisterListServiceHandler)
}
