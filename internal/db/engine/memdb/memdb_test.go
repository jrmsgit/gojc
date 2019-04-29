// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package memdb

import (
	"testing"

	"github.com/jrmsdev/gojc/db/dberr"
	"github.com/jrmsdev/gojc/internal/db/engine"
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
	err := eng.Set("testing", "data")
	check(t, "set error", err, nil)
	err = eng.Set("testing", "data")
	check(t, "is KeyExists error", dberr.Is("KeyExists", err), true)
}

func TestGet(t *testing.T) {
	eng := newEngine(t)
	defer eng.Close()
	err := eng.Set("testing", "data")
	if err != nil {
		t.Fatal(err)
	}
	val := eng.Get("testing")
	check(t, "get value", val, "data")
	val = eng.Get("nokey")
	check(t, "get unset key value", val, "")
}

func TestUpdate(t *testing.T) {
	eng := newEngine(t)
	defer eng.Close()
	err := eng.Set("testing", "data")
	if err != nil {
		t.Fatal(err)
	}
	err = eng.Update("testing", "newdata")
	check(t, "update error", err, nil)
	val := eng.Get("testing")
	check(t, "updated value", val, "newdata")
	err = eng.Update("test", "data")
	check(t, "is KeyNotFound error", dberr.Is("KeyNotFound", err), true)
}

func TestDelete(t *testing.T) {
	eng := newEngine(t)
	defer eng.Close()
	err := eng.Set("testing", "data")
	if err != nil {
		t.Fatal(err)
	}
	err = eng.Delete("testing")
	check(t, "delete error", err, nil)
	err = eng.Delete("testing")
	t.Log(err)
	check(t, "is KeyNotFound error", dberr.Is("KeyNotFound", err), true)
}
