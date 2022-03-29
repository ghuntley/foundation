// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package oci

import (
	"context"
	"fmt"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/empty"
	"github.com/google/go-containerregistry/pkg/v1/mutate"
	"github.com/google/go-containerregistry/pkg/v1/types"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/workspace/compute"
	"namespacelabs.dev/foundation/workspace/tasks"
)

type ImageWithPlatform struct {
	Image    compute.Computable[Image]
	Platform specs.Platform
}

func MakeImageIndex(images ...ImageWithPlatform) compute.Computable[ResolvableImage] {
	return &makeImageIndex{images: images}
}

type makeImageIndex struct {
	images []ImageWithPlatform

	compute.LocalScoped[ResolvableImage]
}

func (al *makeImageIndex) Inputs() *compute.In {
	var platforms []specs.Platform
	in := compute.Inputs()
	for k, d := range al.images {
		in = in.Computable(fmt.Sprintf("image%d", k), d.Image)
		platforms = append(platforms, d.Platform)
	}
	return in.JSON("platforms", platforms)
}

func (al *makeImageIndex) Action() *tasks.ActionEvent {
	var refs []string
	var platforms []specs.Platform
	for _, d := range al.images {
		refs = append(refs, RefFrom(d.Image))
		platforms = append(platforms, d.Platform)
	}
	return tasks.Action("oci.make-image-index").Arg("refs", refs).Arg("platforms", platforms)
}

func (al *makeImageIndex) Compute(ctx context.Context, deps compute.Resolved) (ResolvableImage, error) {
	var adds []mutate.IndexAddendum
	for k, d := range al.images {
		image := compute.GetDepValue(deps, d.Image, fmt.Sprintf("image%d", k))

		digest, err := image.Digest()
		if err != nil {
			return nil, err
		}

		mediaType, err := image.MediaType()
		if err != nil {
			return nil, err
		}

		if mediaType != types.DockerManifestSchema2 {
			return nil, fnerrors.InternalError("%s: unexpected media type: %s", digest.String(), mediaType)
		}

		adds = append(adds, mutate.IndexAddendum{
			Add: image,
			Descriptor: v1.Descriptor{
				MediaType: mediaType,
				Platform: &v1.Platform{
					OS:           d.Platform.OS,
					Architecture: d.Platform.Architecture,
					Variant:      d.Platform.Variant,
				},
			},
		})
	}

	idx := mutate.AppendManifests(mutate.IndexMediaType(empty.Index, types.DockerManifestList), adds...)

	// The Digest() is requested here to guarantee that the index can indeed be created.
	// This will also mark the digest "computed", which is the closest we can get to a
	// sealed result.
	if _, err := idx.Digest(); err != nil {
		return nil, fnerrors.InternalError("failed to compute image index digest: %w", err)
	}

	return rawImageIndex{idx}, nil
}