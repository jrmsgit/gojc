// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package db

import (
	"github.com/jrmsdev/gojc/db/schema"
)

func (d *DB) Get(key string) string {
	q := schema.Query(d.Schema, key)
	if d.Failed() {
		return ""
	}
	return d.eng.Get(q)
}
