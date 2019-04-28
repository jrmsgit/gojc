// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package db

import (
	"github.com/jrmsdev/gojc/db/dberr"
	"github.com/jrmsdev/gojc/internal/db/engine"
	"github.com/jrmsdev/gojc/internal/db/uri"
)

type DB struct {
	eng engine.Engine
	uri *uri.URI
}

func (d *DB) Failed() bool {
	return dberr.Last() != nil
}

func (d *DB) Error() error {
	return dberr.Last()
}
