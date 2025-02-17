// This file was automatically generated by Namespace.
// DO NOT EDIT. To update, re-run `ns generate`.

package postgres

import (
	"context"
	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/std/monitoring/tracing"
)

// Dependencies that are instantiated once for the lifetime of the extension.
type ExtensionDeps struct {
	OpenTelemetry tracing.DeferredTracerProvider
}

type _checkProvideDatabase func(context.Context, *DatabaseArgs, ExtensionDeps) (*DB, error)

var _ _checkProvideDatabase = ProvideDatabase

var (
	Package__sfr1nt = &core.Package{
		PackageName: "namespacelabs.dev/foundation/universe/db/postgres",
	}

	Provider__sfr1nt = core.Provider{
		Package:     Package__sfr1nt,
		Instantiate: makeDeps__sfr1nt,
	}
)

func makeDeps__sfr1nt(ctx context.Context, di core.Dependencies) (_ interface{}, err error) {
	var deps ExtensionDeps

	if err := di.Instantiate(ctx, tracing.Provider__70o2mm, func(ctx context.Context, v interface{}) (err error) {
		if deps.OpenTelemetry, err = tracing.ProvideTracerProvider(ctx, nil, v.(tracing.ExtensionDeps)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return deps, nil
}
