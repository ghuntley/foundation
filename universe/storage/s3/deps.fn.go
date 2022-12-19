// This file was automatically generated by Namespace.
// DO NOT EDIT. To update, re-run `ns generate`.

package s3

import (
	"context"
	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/universe/aws/client"
	"namespacelabs.dev/foundation/universe/aws/s3"
	"namespacelabs.dev/foundation/universe/storage/minio/creds"
)

// Dependencies that are instantiated once for the lifetime of the extension.
type ExtensionDeps struct {
	ClientFactory client.ClientFactory
	MinioCreds    *creds.Creds
}

type _checkProvideBucket func(context.Context, *BucketArgs, ExtensionDeps) (*s3.Bucket, error)

var _ _checkProvideBucket = ProvideBucket

var (
	Package__4pkegf = &core.Package{
		PackageName: "namespacelabs.dev/foundation/universe/storage/s3",
	}

	Provider__4pkegf = core.Provider{
		Package:     Package__4pkegf,
		Instantiate: makeDeps__4pkegf,
	}
)

func makeDeps__4pkegf(ctx context.Context, di core.Dependencies) (_ interface{}, err error) {
	var deps ExtensionDeps

	if err := di.Instantiate(ctx, client.Provider__hva50k, func(ctx context.Context, v interface{}) (err error) {
		if deps.ClientFactory, err = client.ProvideClientFactory(ctx, nil, v.(client.ExtensionDeps)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	if deps.MinioCreds, err = creds.ProvideCreds(ctx, nil); err != nil {
		return nil, err
	}

	return deps, nil
}
