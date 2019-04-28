// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package engine

import (
	"github.com/jrmsdev/gojc/internal/db/uri"
)

type Engine interface {
	Open(*uri.URI) error
}
