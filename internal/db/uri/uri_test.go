// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package uri

import (
	"testing"
)

func check(t *testing.T, name string, got, expect interface{}) {
	t.Helper()
	if got != expect {
		t.Errorf("%s got='%v' expect='%v'", name, got, expect)
	}
}

func TestParse(t *testing.T) {
	u, err := Parse("driver:/dbname")
	if err != nil {
		t.Fatal(err)
	}
	check(t, "db driver", u.Driver, "driver")
	check(t, "db name", u.DBName, "/dbname")
	check(t, "db host", u.Host, "")
}
