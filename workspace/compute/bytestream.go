// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package compute

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"hash"
	"io"
	"os"
	"sync"

	"namespacelabs.dev/foundation/internal/fntypes"
	"namespacelabs.dev/foundation/workspace/dirs"
	"namespacelabs.dev/foundation/workspace/tasks"
)

type ByteStream interface {
	Digest() fntypes.Digest
	ContentLength() uint64
	Reader() (io.ReadCloser, error)
}

func NewByteStream(ctx context.Context) (*ByteStreamWriter, error) {
	f, err := dirs.CreateUserTemp("compute", "bytestream")
	if err != nil {
		return nil, err
	}

	h := sha256.New()
	bsw := &ByteStreamWriter{f: f, h: h, w: io.MultiWriter(h, f), result: &byteStream{path: f.Name()}}

	On(ctx).Cleanup(tasks.Action("compute.output.cleanup"), func(ctx context.Context) error {
		bsw.result.mu.Lock()
		if !bsw.result.consumed {
			os.Remove(bsw.result.path)
		}
		bsw.result.mu.Unlock()
		return nil
	})

	return bsw, nil
}

type byteStream struct {
	path          string
	digest        fntypes.Digest
	contentLength uint64

	mu       sync.Mutex
	consumed bool
}

func (bsw *byteStream) Digest() fntypes.Digest {
	return bsw.digest
}

func (bsw *byteStream) ContentLength() uint64 {
	return bsw.contentLength
}

func (bsw *byteStream) Reader() (io.ReadCloser, error) {
	f, err := os.Open(bsw.path)
	return f, err
}

var _ json.Marshaler = &byteStream{}

func (bsw *byteStream) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"path":          bsw.path,
		"digest":        bsw.digest,
		"contentLength": bsw.contentLength,
	})
}

type ByteStreamWriter struct {
	f         *os.File
	h         hash.Hash
	w         io.Writer
	byteCount uint64
	result    *byteStream
}

var _ io.WriteCloser = &ByteStreamWriter{}

func (bsw *ByteStreamWriter) Write(p []byte) (int, error) {
	n, err := bsw.w.Write(p)
	bsw.byteCount += uint64(n)
	return n, err
}

func (bsw *ByteStreamWriter) Close() error {
	return bsw.f.Close()
}

func (bsw *ByteStreamWriter) Complete() (ByteStream, error) {
	if err := bsw.f.Close(); err != nil {
		return nil, err
	}

	d := fntypes.FromHash("sha256", bsw.h)
	bsw.result.digest = d
	bsw.result.contentLength = bsw.byteCount
	return bsw.result, nil
}