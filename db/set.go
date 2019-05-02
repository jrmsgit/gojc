// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package db

import (
	"github.com/jrmsdev/gojc/db/schema"
)

func (d *DB) Set(key, val string) error {
	stmt := schema.Stmt(d.Schema, key)
	if d.Failed() {
		return d.Error()
	}
	return d.eng.Set(stmt, val)
}
