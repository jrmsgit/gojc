// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package engine

import (
	"github.com/jrmsdev/gojc/internal/db/uri"
)

type Engine interface {
	Open(*uri.URI) error
}

type EngineCreator func() Engine

var reg map[string]EngineCreator

func init() {
	reg = make(map[string]EngineCreator)
}

func Register(creator EngineCreator, driver ...string) {
	for _, drv := range driver {
		reg[drv] = creator
	}
}
