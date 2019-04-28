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

var uriTests = map[string][]string{
	// uri: {driver, db, host, port, username, password}
	"driver:/dbname": {"driver", "/dbname", "", "", "", ""},
}

func TestParse(t *testing.T) {
	for rawuri, x := range uriTests {
		u, err := Parse(rawuri)
		if err != nil {
			t.Fatal(err)
		}
		check(t, "db driver", u.Driver, x[0])
		check(t, "db name", u.DBName, x[1])
		check(t, "db host", u.Host, x[2])
		check(t, "db port", u.Port, x[3])
		check(t, "db username", u.Username, x[4])
		check(t, "db password", u.GetPassword(), x[5])
	}
}

func TestParserError(t *testing.T) {
	_, err := Parse("::")
	check(t, "parse error", err.Error(), "parse ::: missing protocol scheme")
}

func TestNoDriver(t *testing.T) {
	_, err := Parse("")
	check(t, "parse error", err.Error(), "db driver not set")
}

func TestNoDBName(t *testing.T) {
	_, err := Parse("memdb:")
	check(t, "parse error", err.Error(), "db name not set")
}
