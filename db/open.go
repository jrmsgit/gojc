// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package db

import (
	"github.com/jrmsdev/gojc/internal/db/engine"
	"github.com/jrmsdev/gojc/internal/db/uri"
)

func Open(rawuri string) (*DB, error) {
	d := new(DB)
	if u, err := uri.Parse(rawuri); err != nil {
		return nil, err
	} else {
		d.uri = u
	}
	if eng, err := engine.Get(d.uri); err != nil {
		return nil, err
	} else {
		d.eng = eng
	}
	if err := d.eng.Open(d.uri); err != nil {
		return nil, err
	}
	return d, nil
}
