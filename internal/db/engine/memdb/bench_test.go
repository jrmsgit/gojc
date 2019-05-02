// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package memdb

import (
	"strconv"
	"testing"

	"github.com/jrmsdev/gojc/internal/db/engine"
	"github.com/jrmsdev/gojc/internal/db/query"
	"github.com/jrmsdev/gojc/internal/db/statement"
	"github.com/jrmsdev/gojc/internal/db/uri"
)

func benchEngine(b *testing.B) engine.Engine {
	b.Helper()
	u, err := uri.Parse("db:/test")
	if err != nil {
		b.Fatal(err)
	}
	eng := New()
	err = eng.Open(u)
	if err != nil {
		b.Fatal(err)
	}
	return eng
}

func BenchmarkSet(b *testing.B) {
	eng := benchEngine(b)
	defer eng.Close()
	for i := 0; i < b.N; i++ {
		stmt := statement.New(strconv.Itoa(i))
		eng.Set(stmt, "data")
	}
}

func BenchmarkGet(b *testing.B) {
	eng := benchEngine(b)
	defer eng.Close()
	stmt := statement.New("testing")
	eng.Set(stmt, "data")
	q := query.New("testing")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		eng.Get(q)
	}
}

func BenchmarkGetAll(b *testing.B) {
	eng := benchEngine(b)
	defer eng.Close()
	stmt := statement.New("test.0")
	eng.Set(stmt, "data0")
	stmt = statement.New("test.1")
	eng.Set(stmt, "data1")
	stmt = statement.New("test.2")
	eng.Set(stmt, "data2")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		eng.GetAll("testing")
	}
}

func BenchmarkUpdate(b *testing.B) {
	eng := benchEngine(b)
	defer eng.Close()
	stmt := statement.New("testing")
	eng.Set(stmt, "data")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		eng.Update("testing", "newdata")
	}
}

func BenchmarkDelete(b *testing.B) {
	eng := benchEngine(b)
	defer eng.Close()
	for i := 0; i < b.N; i++ {
		stmt := statement.New(strconv.Itoa(i))
		eng.Set(stmt, "data")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		eng.Delete(strconv.Itoa(i))
	}
}
