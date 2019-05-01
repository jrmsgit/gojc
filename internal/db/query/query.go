// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package query

type Query struct {
	key string
}

func New(key string) *Query {
	q := new(Query)
	q.key = key
	return q
}

func (q *Query) Key() string {
	return q.key
}
