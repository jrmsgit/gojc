// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package schema

import (
	"github.com/jrmsdev/gojc/internal/db/statement"
)

func Stmt(s *Schema, key string) *statement.Stmt {
	return statement.New(key)
}
