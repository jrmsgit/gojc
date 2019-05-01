// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package db

func (d *DB) Update(key, val string) error {
	return d.eng.Update(key, val)
}
