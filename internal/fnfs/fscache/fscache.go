// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package fscache

import (
	"context"
	"io"
	"io/fs"
	"reflect"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"namespacelabs.dev/foundation/internal/artifacts/oci"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/fnfs/tarfs"
	"namespacelabs.dev/foundation/internal/fntypes"
	"namespacelabs.dev/foundation/workspace/cache"
	"namespacelabs.dev/foundation/workspace/compute"
)

func RegisterFSCacheable() {
	compute.RegisterCacheable[fs.FS](fsCacheable{})
}

// Implementations of a `fs.FS` cache where each FS is cached as a individual image layer.
type fsCacheable struct{}

func ComputeDigest(ctx context.Context, fsys fs.FS) (fntypes.Digest, error) {
	layer, err := oci.LayerFromFS(ctx, fsys)
	if err != nil {
		return fntypes.Digest{}, err
	}

	_, d, err := oci.ComputeLayerCacheData(layer)
	return d, err
}

func (fsCacheable) ComputeDigest(ctx context.Context, v interface{}) (fntypes.Digest, error) {
	return ComputeDigest(ctx, v.(fs.FS)) // XXX don't pay the cost twice (see Cache below).
}

func (fsCacheable) LoadCached(ctx context.Context, c cache.Cache, t reflect.Type, h fntypes.Digest) (compute.Result[fs.FS], error) {
	layer, digest, err := oci.LoadCachedLayer(ctx, c, h)
	if err != nil {
		return compute.Result[fs.FS]{}, err
	}

	return compute.Result[fs.FS]{
		Digest: digest,
		Value:  layerBackedFS{layer},
	}, nil
}

func (fsCacheable) Cache(ctx context.Context, c cache.Cache, vfs fs.FS) (fntypes.Digest, error) {
	layer, err := oci.LayerFromFS(ctx, vfs)
	if err != nil {
		return fntypes.Digest{}, err
	}

	return oci.CacheLayer(ctx, c, layer)
}

// Implements a fs.ReadDirFS which is backed by a layer. We don't buffer the layer in memory though,
// its contents are read on demand.
type layerBackedFS struct{ l v1.Layer }

var _ fs.ReadDirFS = layerBackedFS{}
var _ oci.HasToLayer = layerBackedFS{}

func (l layerBackedFS) tarStream() (io.ReadCloser, error) {
	digest, err := l.l.Digest()
	if err != nil {
		return nil, fnerrors.BadInputError("failed to get layer digest: %w", err)
	}

	r, err := l.l.Uncompressed()
	if err != nil {
		return nil, fnerrors.BadInputError("%s: failed to get layer contents", digest)
	}

	return r, nil
}

func (l layerBackedFS) Open(path string) (fs.File, error) {
	return tarfs.FS{TarStream: l.tarStream}.Open(path)
}

func (l layerBackedFS) ReadDir(dir string) ([]fs.DirEntry, error) {
	return tarfs.FS{TarStream: l.tarStream}.ReadDir(dir)
}

func (l layerBackedFS) VisitFiles(ctx context.Context, visitor func(string, []byte, fs.DirEntry) error) error {
	return tarfs.FS{TarStream: l.tarStream}.VisitFiles(ctx, visitor)
}

func (f layerBackedFS) AsLayer() (v1.Layer, error) { return f.l, nil }