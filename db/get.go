// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package db

func (d *DB) Get(key string) string {
	return d.eng.Get(key)
}
