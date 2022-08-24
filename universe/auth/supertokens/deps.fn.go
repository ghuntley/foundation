// This file was automatically generated by Namespace.
// DO NOT EDIT. To update, re-run `ns generate`.

package supertokens

import (
	"context"
	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/std/go/http/middleware"
	"namespacelabs.dev/foundation/std/secrets"
)

// Dependencies that are instantiated once for the lifetime of the extension.
type ExtensionDeps struct {
	GithubClientId     *secrets.Value
	GithubClientSecret *secrets.Value
	Middleware         middleware.Middleware
}

var (
	Package__q3mscu = &core.Package{
		PackageName: "namespacelabs.dev/foundation/universe/auth/supertokens",
	}

	Provider__q3mscu = core.Provider{
		Package:     Package__q3mscu,
		Instantiate: makeDeps__q3mscu,
	}

	Initializers__q3mscu = []*core.Initializer{
		{
			Package: Package__q3mscu,
			Do: func(ctx context.Context, di core.Dependencies) error {
				return di.Instantiate(ctx, Provider__q3mscu, func(ctx context.Context, v interface{}) error {
					return Prepare(ctx, v.(ExtensionDeps))
				})
			},
		},
	}
)

func makeDeps__q3mscu(ctx context.Context, di core.Dependencies) (_ interface{}, err error) {
	var deps ExtensionDeps

	// name: "github_client_id"
	if deps.GithubClientId, err = secrets.ProvideSecret(ctx, core.MustUnwrapProto("ChBnaXRodWJfY2xpZW50X2lk", &secrets.Secret{}).(*secrets.Secret)); err != nil {
		return nil, err
	}

	// name: "github_client_secret"
	if deps.GithubClientSecret, err = secrets.ProvideSecret(ctx, core.MustUnwrapProto("ChRnaXRodWJfY2xpZW50X3NlY3JldA==", &secrets.Secret{}).(*secrets.Secret)); err != nil {
		return nil, err
	}

	if deps.Middleware, err = middleware.ProvideMiddleware(ctx, nil); err != nil {
		return nil, err
	}

	return deps, nil
}
