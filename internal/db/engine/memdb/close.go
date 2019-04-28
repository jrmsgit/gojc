// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package memdb

import (
	"github.com/jrmsdev/gojc/db/dberr"
)

func (d *DB) Close() error {
	if d.data == nil {
		return dberr.Set("NotOpen")
	}
	d.data = nil
	return nil
}
