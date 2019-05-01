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
	"NotOpen":  errors.New("db is not open"),
}

var setreg = map[string]error{
	"UriParse":      nil,
	"InvalidDriver": nil,
	"KeyExists":     nil,
	"KeyNotFound":   nil,
	"SchemaError":   nil,
}

func Last() error {
	return last
}

func Set(typ string) error {
	e, ok := reg[typ]
	if !ok {
		return errors.New("InvalidErrorType:" + typ)
	}
	last = e
	return last
}

func SetError(typ string, format string, args ...interface{}) error {
	_, ok := setreg[typ]
	if !ok {
		return errors.New("SetInvalidErrorType:" + typ)
	}
	last = errors.New(fmt.Sprintf(format, args...))
	setreg[typ] = last
	return last
}

func Clear() {
	last = nil
}

func Is(typ string, err error) bool {
	x, ok := reg[typ]
	if ok {
		return err == x
	}
	x, ok = setreg[typ]
	if !ok {
		x = errors.New("IsInvalidErrorType:" + typ)
	}
	return err == x
}
