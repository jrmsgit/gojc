// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package memdb

import (
	"github.com/jrmsdev/gojc/db/dberr"
	"github.com/jrmsdev/gojc/internal/db/statement"
)

func (d *DB) Set(stmt *statement.Stmt, val string) error {
	key := stmt.Key()
	if _, exists := d.data[key]; exists {
		return dberr.SetError("KeyExists", "db key '%s' already exists", key)
	}
	d.data[key] = val
	return nil
}
