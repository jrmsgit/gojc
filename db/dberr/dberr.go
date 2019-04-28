// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package dberr

import (
	"errors"
	"fmt"
)

var last error
var reg = map[string]error{
	"NoDriver": errors.New("db driver not set"),
	"NoDBName": errors.New("db name not set"),
	"IsOpen":   errors.New("db already open"),
	"UriParse": nil,
	"InvalidDriver": nil,
}

func get(typ string) error {
	e, ok := reg[typ]
	if !ok {
		return errors.New("InvalidErrorType:" + typ)
	}
	return e
}

func Last() error {
	return last
}

func Set(typ string) error {
	last = get(typ)
	return last
}

func SetError(typ string, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	err := errors.New(msg)
	if get(typ) == nil {
		last = err
		reg[typ] = last
	} else {
		e := errors.New("SetInvalidErrorType:" + typ)
		last = e
	}
	return last
}

func Clear() {
	last = nil
}

func Is(typ string, err error) bool {
	return err == get(typ)
}
