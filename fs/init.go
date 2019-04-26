// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
)

var storage map[string]*zip.File
var b64 = base64.StdEncoding.DecodeString

func Init(b64zip string) error {
	storage = make(map[string]*zip.File)
	if b64zip != "" {
		blob, err := b64(b64zip)
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
