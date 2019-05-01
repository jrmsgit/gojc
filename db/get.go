// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package db

import (
	"github.com/jrmsdev/gojc/db/dberr"
	"github.com/jrmsdev/gojc/db/schema"
)

func (d *DB) Get(key string) string {
	q, err := schema.Query(d.Schema, key)
	if err != nil {
		dberr.SetError("SchemaError", "%s", err)
		return ""
	}
	return d.eng.Get(q)
}
