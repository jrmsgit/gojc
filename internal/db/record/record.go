// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package record

type Record struct {
}

func New(s string) (*Record, error) {
	r := new(Record)
	return r, nil
}
