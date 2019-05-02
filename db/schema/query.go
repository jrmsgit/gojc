// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package schema

import (
	"github.com/jrmsdev/gojc/internal/db/query"
)

func Query(s *Schema, key string) *query.Query {
	return query.New(key)
}
