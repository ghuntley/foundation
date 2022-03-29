// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package ctxio

import (
	"context"
	"io"
)

type writer struct {
	w        io.Writer
	ctx      context.Context
	progress func(int64)
}

func WriterWithContext(ctx context.Context, w io.Writer, progress func(int64)) io.Writer {
	return writer{
		w:        w,
		ctx:      ctx,
		progress: progress,
	}
}

func (w writer) Write(p []byte) (int, error) {
	if w.ctx.Err() != nil {
		return 0, w.ctx.Err()
	}

	n, err := w.w.Write(p)

	if w.progress != nil {
		w.progress(int64(n))
	}

	return n, err
}