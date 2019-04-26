// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"io"
	"os"
)

func Open(filename string) (io.ReadCloser, error) {
	fh, err := os.Open(filename)
	if err != nil {
		if zipExists(filename) {
			return zipOpen(filename)
		}
		return nil, err
	}
	return fh, nil
}

func zipExists(filename string) bool {
	_, ok := storage[filename]
	return ok
}

func zipOpen(filename string) (io.ReadCloser, error) {
	fh := storage[filename]
	return fh.Open()
}
