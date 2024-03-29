// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package memdb

import (
	"github.com/jrmsdev/gojc/internal/db/query"
)

func (d *DB) Get(q *query.Query) string {
	return d.data[q.Key()]
}
