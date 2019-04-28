// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package dberr

import (
	"errors"
)

var NoDriver = errors.New("db driver not set")
var NoDBName = errors.New("db name not set")
var IsOpen = errors.New("db already open")
