// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package memdb

import (
	"errors"

	"github.com/jrmsdev/gojc/internal/db/record"
	"github.com/jrmsdev/gojc/internal/db/uri"
)

func (d *DB) Open(u *uri.URI) error {
	if d.data != nil {
		return errors.New("db already open")
	}
	d.data = make(map[string]*record.Record)
	return nil
}
