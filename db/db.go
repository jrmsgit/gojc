// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package db

import (
	"github.com/jrmsdev/gojc/db/dberr"
	"github.com/jrmsdev/gojc/internal/db/engine"
	"github.com/jrmsdev/gojc/internal/db/uri"
)

type DB struct {
	uri *uri.URI
	eng engine.Engine
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

func (d *DB) Close() error {
	return d.eng.Close()
}

func (d *DB) Get(key string) string {
	return d.eng.Get(key)
}

func (d *DB) Set(key, val string) error {
	return d.eng.Set(key, val)
}
