// This file was automatically generated by Namespace.
// DO NOT EDIT. To update, re-run `ns generate`.

package honeycomb

import (
	"context"
	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/std/monitoring/tracing"
)

// Dependencies that are instantiated once for the lifetime of the extension.
type ExtensionDeps struct {
	OpenTelemetry tracing.Exporter
}

var (
	Package__e1uscp = &core.Package{
		PackageName: "namespacelabs.dev/foundation/universe/monitoring/honeycomb",
	}

	Provider__e1uscp = core.Provider{
		Package:     Package__e1uscp,
		Instantiate: makeDeps__e1uscp,
	}

	Initializers__e1uscp = []*core.Initializer{
		{
			Package: Package__e1uscp,
			Before:  []string{"namespacelabs.dev/foundation/std/monitoring/tracing"},
			Do: func(ctx context.Context, di core.Dependencies) error {
				return di.Instantiate(ctx, Provider__e1uscp, func(ctx context.Context, v interface{}) error {
					return Prepare(ctx, v.(ExtensionDeps))
				})
			},
		},
	}
)

func makeDeps__e1uscp(ctx context.Context, di core.Dependencies) (_ interface{}, err error) {
	var deps ExtensionDeps

	if err := di.Instantiate(ctx, tracing.Provider__70o2mm, func(ctx context.Context, v interface{}) (err error) {
		// name: "honeycomb"
		if deps.OpenTelemetry, err = tracing.ProvideExporter(ctx, core.MustUnwrapProto("Cglob25leWNvbWI=", &tracing.ExporterArgs{}).(*tracing.ExporterArgs), v.(tracing.ExtensionDeps)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return deps, nil
}
