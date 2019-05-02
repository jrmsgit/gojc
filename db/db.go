// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package db

import (
	"github.com/jrmsdev/gojc/db/dberr"
	"github.com/jrmsdev/gojc/db/schema"
	"github.com/jrmsdev/gojc/internal/db/engine"
	"github.com/jrmsdev/gojc/internal/db/uri"
)

type DB struct {
	uri    *uri.URI
	eng    engine.Engine
	Schema *schema.Schema
}

func New(rawuri string) (*DB, error) {
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
	d.Schema = schema.New()
	return d, nil
}

func (d *DB) Name() string {
	return d.uri.DBName
}

func (d *DB) Failed() bool {
	return dberr.Last() != nil
}

func (d *DB) Error() error {
	return dberr.Last()
}
