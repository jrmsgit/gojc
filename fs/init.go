// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fs

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"errors"
)

var prefix string
var storage map[string]map[string]*zip.File
var b64 = base64.StdEncoding.DecodeString

func init() {
	storage = make(map[string]map[string]*zip.File)
}

func Init(prefix, zipfile string) error {
	if _, done := storage[prefix]; done {
		return errors.New("zip fs init already done for prefix " + prefix)
	}
	storage[prefix] = make(map[string]*zip.File)
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
			storage[prefix][f.Name] = f
		}
	}
	return nil
}
