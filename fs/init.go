// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
)

var prefix string
var storage map[string]*zip.File
var b64 = base64.StdEncoding.DecodeString

func Init(prefix, zipfile string) error {
	storage = make(map[string]*zip.File)
	if zipfile != "" {
		blob, err := b64(zipfile)
		if err != nil {
			return err
		}
		zdata := bytes.NewReader(blob)
		var zr *zip.Reader
		zr, err = zip.NewReader(zdata, int64(zdata.Len()))
		if err != nil {
			return err
		}
		for _, f := range zr.File {
			storage[f.Name] = f
		}
	}
	return nil
}
