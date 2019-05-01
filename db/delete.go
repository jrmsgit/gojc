// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package db

func (d *DB) Delete(key string) error {
	return d.eng.Delete(key)
}
