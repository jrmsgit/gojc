// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package memdb

import (
	"github.com/jrmsdev/gojc/db/dberr"
)

func (d *DB) Set(key, val string) error {
	_, exists := d.data[key]
	if exists {
		return dberr.SetError("KeyExists", "db key '%s' already exists", key)
	}
	d.data[key] = val
	return nil
}
