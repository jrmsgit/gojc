// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package uri

import (
	"testing"

	"github.com/jrmsdev/gojc/db/dberr"
)

func check(t *testing.T, name string, got, expect interface{}) {
	t.Helper()
	if got != expect {
		t.Errorf("%s got='%v' expect='%v'", name, got, expect)
	}
}

var uriTests = map[string][]string{
	"driver:/dbname":             {"driver", "/dbname", "", "", "", ""},
	"drv://host/db":              {"drv", "/db", "host", "", "", ""},
	"drv://@host/db":             {"drv", "/db", "host", "", "", ""},
	"drv://host:port/db":         {"drv", "/db", "host", "port", "", ""},
	"drv://user@host/db":         {"drv", "/db", "host", "", "user", ""},
	"drv://user:pw@host/db":      {"drv", "/db", "host", "", "user", "pw"},
	"drv://user:@host/db":        {"drv", "/db", "host", "", "user", ""},
	"drv://:pw@host/db":          {"drv", "/db", "host", "", "", "pw"},
	"drv://:pw@host:port/db":     {"drv", "/db", "host", "port", "", "pw"},
	"drv://user:pw@host:port/db": {"drv", "/db", "host", "port", "user", "pw"},
}

func TestParse(t *testing.T) {
	for rawuri, x := range uriTests {
		u, err := Parse(rawuri)
		if err != nil {
			t.Fatalf("error uri='%s' %s", rawuri, err)
		}
		check(t, "db driver", u.Driver, x[0])
		check(t, "db name", u.DBName, x[1])
		check(t, "db host", u.Host, x[2])
		check(t, "db port", u.Port, x[3])
		check(t, "db username", u.Username, x[4])
		check(t, "db password", u.Password, x[5])
	}
}

func TestParserError(t *testing.T) {
	_, err := Parse("::")
	check(t, "parse error", err.Error(), "parse ::: missing protocol scheme")
}

func TestNoDriver(t *testing.T) {
	_, err := Parse("")
	check(t, "parse error", err, dberr.NoDriver)
}

func TestNoDBName(t *testing.T) {
	_, err := Parse("driver:")
	check(t, "parse error", err, dberr.NoDBName)
}
