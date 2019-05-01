// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package db

func (d *DB) GetAll(key string) map[string]string {
	return d.eng.GetAll(key)
}
