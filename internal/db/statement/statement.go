// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package statement

type Stmt struct {
	key string
}

func New(key string) *Stmt {
	s := new(Stmt)
	s.key = key
	return s
}

func (s *Stmt) Key() string {
	return s.key
}
