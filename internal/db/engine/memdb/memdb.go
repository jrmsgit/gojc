// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package memdb

import (
	"github.com/jrmsdev/gojc/internal/db/engine"
	"github.com/jrmsdev/gojc/internal/db/record"
	"github.com/jrmsdev/gojc/internal/db/uri"
)

func init() {
	engine.Register(New, "memdb", "db")
}

type DB struct {
	data map[string]*record.Record
}

func New() engine.Engine {
	return new(DB)
}

func (d *DB) Open(u *uri.URI) error {
	d.data = make(map[string]*record.Record)
	return nil
}
