// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package engine_test

import (
	"testing"

	"github.com/jrmsdev/gojc/db/dberr"
	"github.com/jrmsdev/gojc/internal/db/engine"
	"github.com/jrmsdev/gojc/internal/db/uri"

	_ "github.com/jrmsdev/gojc/internal/db/engine/memdb"
)

func check(t *testing.T, name string, got, expect interface{}) {
	t.Helper()
	if got != expect {
		t.Errorf("%s got='%v' expect='%v'", name, got, expect)
	}
}

func TestRegister(t *testing.T) {
	check(t, "number of drivers", len(engine.Drivers()), 2)
	check(t, "db driver", engine.HasDriver("db"), true)
	check(t, "memdb driver", engine.HasDriver("memdb"), true)
}

func TestGetEngine(t *testing.T) {
	u, uerr := uri.Parse("db:/test")
	if uerr != nil {
		t.Fatal(uerr)
	}
	_, err := engine.Get(u)
	check(t, "get engine error", err, nil)
}

func TestGetError(t *testing.T) {
	u, uerr := uri.Parse("drv:/test")
	if uerr != nil {
		t.Fatal(uerr)
	}
	_, err := engine.Get(u)
	check(t, "InvalidDriver error", dberr.Is("InvalidDriver", err), true)
}
