// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package uri

import (
	"net/url"

	"github.com/jrmsdev/gojc/db/dberr"
)

type URI struct {
	Driver   string
	DBName   string
	Host     string
	Port     string
	Username string
	Password string
}

func Parse(rawuri string) (*URI, error) {
	u, err := url.Parse(rawuri)
	if err != nil {
		return nil, dberr.SetError("UriParse", "%s", err)
	}
	if err := checkArgs(u); err != nil {
		return nil, err
	}
	r := new(URI)
	r.Driver = u.Scheme
	r.DBName = u.EscapedPath()
	r.Host = u.Hostname()
	r.Port = u.Port()
	r.Username = u.User.Username()
	r.Password, _ = u.User.Password()
	return r, nil
}

func checkArgs(u *url.URL) error {
	if u.Scheme == "" {
		return dberr.Set("NoDriver")
	}
	if u.Path == "" {
		return dberr.Set("NoDBName")
	}
	return nil
}
