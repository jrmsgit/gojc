// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package engine

import (
	"github.com/jrmsdev/gojc/db/dberr"
	"github.com/jrmsdev/gojc/internal/db/query"
	"github.com/jrmsdev/gojc/internal/db/statement"
	"github.com/jrmsdev/gojc/internal/db/uri"
)

type Engine interface {
	Open(*uri.URI) error
	Close() error
	Get(q *query.Query) string
	GetAll(key string) map[string]string
	Set(stmt *statement.Stmt, val string) error
	Update(key, val string) error
	Delete(key string) error
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

func Drivers() []string {
	l := make([]string, 0)
	for n := range reg {
		l = append(l, n)
	}
	return l
}

func HasDriver(name string) bool {
	_, ok := reg[name]
	return ok
}

func Get(u *uri.URI) (Engine, error) {
	mk, ok := reg[u.Driver]
	if !ok {
		return nil, dberr.SetError("InvalidDriver", "invalid driver %s", u.Driver)
	}
	return mk(), nil
}
