// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"io"
	"os"
	"path/filepath"
)

func Open(prefix, filename string) (io.ReadCloser, error) {
	fh, err := os.Open(filepath.Join(prefix, filename))
	if err != nil {
		if zipExists(prefix, filename) {
			return zipOpen(prefix, filename)
		}
		return nil, err
	}
	return fh, nil
}

func zipExists(prefix, filename string) bool {
	_, ok := storage[prefix][filename]
	return ok
}

func zipOpen(prefix, filename string) (io.ReadCloser, error) {
	fh := storage[prefix][filename]
	return fh.Open()
}
