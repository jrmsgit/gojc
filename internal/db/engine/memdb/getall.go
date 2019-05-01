// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package memdb

import (
	"strings"
)

func (d *DB) GetAll(key string) map[string]string {
	r := make(map[string]string)
	x := key + "."
	for k, v := range d.data {
		if strings.HasPrefix(k, x) {
			n := strings.Replace(k, x, "", 1)
			r[n] = strings.TrimSpace(v)
		}
	}
	return r
}
