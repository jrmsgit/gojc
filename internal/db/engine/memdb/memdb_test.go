// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package memdb

import (
	"testing"

	"github.com/jrmsdev/gojc/db/dberr"
	"github.com/jrmsdev/gojc/internal/db/engine"
	"github.com/jrmsdev/gojc/internal/db/query"
	"github.com/jrmsdev/gojc/internal/db/statement"
	"github.com/jrmsdev/gojc/internal/db/uri"
)

func check(t *testing.T, name string, got, expect interface{}) {
	t.Helper()
	if got != expect {
		t.Errorf("%s got='%v' expect='%v'", name, got, expect)
	}
}

func TestOpen(t *testing.T) {
	u, err := uri.Parse("db:/test")
	if err != nil {
		t.Fatal(err)
	}
	eng := New()
	err = eng.Open(u)
	check(t, "open error", err, nil)
	err = eng.Open(u)
	check(t, "is IsOpen error", dberr.Is("IsOpen", err), true)
}

func TestClose(t *testing.T) {
	u, err := uri.Parse("db:/test")
	if err != nil {
		t.Fatal(err)
	}
	eng := New()
	err = eng.Open(u)
	if err != nil {
		t.Fatal(err)
	}
	err = eng.Close()
	check(t, "close error", err, nil)
	err = eng.Close()
	check(t, "is NotOpen error", dberr.Is("NotOpen", err), true)
}

func newEngine(t *testing.T) engine.Engine {
	t.Helper()
	u, err := uri.Parse("db:/test")
	if err != nil {
		t.Fatal(err)
	}
	eng := New()
	err = eng.Open(u)
	if err != nil {
		t.Fatal(err)
	}
	return eng
}

func TestSet(t *testing.T) {
	eng := newEngine(t)
	defer eng.Close()
	stmt := statement.New("testing")
	err := eng.Set(stmt, "data")
	check(t, "set error", err, nil)
	err = eng.Set(stmt, "data")
	check(t, "is KeyExists error", dberr.Is("KeyExists", err), true)
}

func TestGet(t *testing.T) {
	eng := newEngine(t)
	defer eng.Close()
	stmt := statement.New("testing")
	err := eng.Set(stmt, "data")
	if err != nil {
		t.Fatal(err)
	}
	q := query.New("testing")
	val := eng.Get(q)
	check(t, "get value", val, "data")
	q = query.New("nokey")
	val = eng.Get(q)
	check(t, "get unset key value", val, "")
}

func TestGetAll(t *testing.T) {
	eng := newEngine(t)
	defer eng.Close()
	stmt := statement.New("testing.f0")
	err := eng.Set(stmt, "v0")
	if err != nil {
		t.Fatal(err)
	}
	stmt = statement.New("testing.f1")
	err = eng.Set(stmt, "v1")
	if err != nil {
		t.Fatal(err)
	}
	stmt = statement.New("test.f2")
	err = eng.Set(stmt, "v2")
	if err != nil {
		t.Fatal(err)
	}
	rst := eng.GetAll("testing")
	check(t, "get f0 value", rst["f0"], "v0")
	check(t, "get f1 value", rst["f1"], "v1")
	check(t, "get f2 value", rst["f2"], "")
	rst = eng.GetAll("test")
	check(t, "get 2nd f0 value", rst["f0"], "")
	check(t, "get 2nd f2 value", rst["f2"], "v2")
}

func TestUpdate(t *testing.T) {
	eng := newEngine(t)
	defer eng.Close()
	stmt := statement.New("testing")
	err := eng.Set(stmt, "data")
	if err != nil {
		t.Fatal(err)
	}
	err = eng.Update("testing", "newdata")
	check(t, "update error", err, nil)
	q := query.New("testing")
	val := eng.Get(q)
	check(t, "updated value", val, "newdata")
	err = eng.Update("test", "data")
	check(t, "is KeyNotFound error", dberr.Is("KeyNotFound", err), true)
}

func TestDelete(t *testing.T) {
	eng := newEngine(t)
	defer eng.Close()
	stmt := statement.New("testing")
	err := eng.Set(stmt, "data")
	if err != nil {
		t.Fatal(err)
	}
	err = eng.Delete("testing")
	check(t, "delete error", err, nil)
	err = eng.Delete("testing")
	check(t, "is KeyNotFound error", dberr.Is("KeyNotFound", err), true)
}
