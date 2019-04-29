// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package memdb

import (
	"github.com/jrmsdev/gojc/db/dberr"
)

func (d *DB) Update(key, val string) error {
	_, exists := d.data[key]
	if !exists {
		return dberr.SetError("KeyNotFound", "db key '%s' not found", key)
	}
	delete(d.data, key)
	d.data[key] = val
	return nil
}
