// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package memdb

func (d *DB) Get(key string) string {
	return d.data[key]
}
