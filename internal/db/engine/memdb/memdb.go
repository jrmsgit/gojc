// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package memdb

import (
	"github.com/jrmsdev/gojc/internal/db/engine"
	"github.com/jrmsdev/gojc/internal/db/record"
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
