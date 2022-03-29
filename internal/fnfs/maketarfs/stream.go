// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package maketarfs

import (
	"archive/tar"
	"context"
	"io"
	"io/fs"
	"path/filepath"
	"time"

	"namespacelabs.dev/foundation/internal/fnfs"
	"namespacelabs.dev/foundation/internal/uniquestrings"
)

var FixedPoint = time.Unix(1, 1)

func TarFS(ctx context.Context, parentW io.Writer, vfs fs.FS, includeFiles []string, excludeFiles []string) error {
	w := tar.NewWriter(parentW)
	defer w.Close()

	var inclusion uniquestrings.List
	for _, f := range includeFiles {
		inclusion.Add(f)
	}

	var exclusion uniquestrings.List
	for _, f := range excludeFiles {
		exclusion.Add(f)
	}

	dirs := map[string]bool{}

	if err := fnfs.VisitFiles(ctx, vfs, func(path string, contents []byte, _ fs.DirEntry) error {
		if exclusion.Has(path) || (len(includeFiles) > 0 && !inclusion.Has(path)) {
			return nil
		}

		dir := filepath.Dir(path)
		if dir != "." && !dirs[dir] {
			if err := w.WriteHeader(&tar.Header{
				Name:     dir,
				Typeflag: tar.TypeDir,
				Mode:     0555,
				ModTime:  FixedPoint,
			}); err != nil {
				return err
			}
			dirs[dir] = true
		}

		if err := w.WriteHeader(&tar.Header{
			Name:     path,
			Size:     int64(len(contents)),
			Typeflag: tar.TypeReg,
			Mode:     0555,
			ModTime:  FixedPoint,
		}); err != nil {
			return err
		}

		if _, err := w.Write(contents); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}

	return nil
}