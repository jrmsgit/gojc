// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package memdb

import (
	"testing"

	"github.com/jrmsdev/gojc/db/dberr"
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
	check(t, "IsOpen error", dberr.Is("IsOpen", err), true)
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
	check(t, "NotOpen error", dberr.Is("NotOpen", err), true)
}
