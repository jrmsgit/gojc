// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package db

func (d *DB) Set(key, val string) error {
	return d.eng.Set(key, val)
}
