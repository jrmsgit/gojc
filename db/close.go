// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package db

func (d *DB) Close() error {
	return d.eng.Close()
}
