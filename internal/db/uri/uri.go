// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package uri

import (
	"errors"
	"net/url"
)

type URI struct {
	Driver string
	DBName string
	Host   string
}

func Parse(rawuri string) (*URI, error) {
	u, err := url.Parse(rawuri)
	if err != nil {
		return nil, err
	}
	if err := checkArgs(u); err != nil {
		return nil, err
	}
	r := new(URI)
	r.Driver = u.Scheme
	r.DBName = u.Path
	r.Host = u.Host
	return r, nil
}

func checkArgs(u *url.URL) error {
	if u.Scheme == "" {
		return errors.New("db driver not set")
	}
	if u.Path == "" {
		return errors.New("db name not set")
	}
	return nil
}
