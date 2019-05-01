// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package memdb

import (
	"testing"
	"strconv"

	"github.com/jrmsdev/gojc/internal/db/engine"
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
		eng.Set(strconv.Itoa(i), "data")
	}
}

func BenchmarkGet(b *testing.B) {
	eng := benchEngine(b)
	defer eng.Close()
	eng.Set("testing", "data")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		eng.Get("testing")
	}
}

func BenchmarkGetAll(b *testing.B) {
	eng := benchEngine(b)
	defer eng.Close()
	eng.Set("test.0", "data0")
	eng.Set("test.1", "data1")
	eng.Set("test.2", "data2")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		eng.GetAll("testing")
	}
}

func BenchmarkUpdate(b *testing.B) {
	eng := benchEngine(b)
	defer eng.Close()
	eng.Set("testing", "data")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		eng.Update("testing", "newdata")
	}
}

func BenchmarkDelete(b *testing.B) {
	eng := benchEngine(b)
	defer eng.Close()
	for i := 0; i < b.N; i++ {
		eng.Set(strconv.Itoa(i), "data")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		eng.Delete(strconv.Itoa(i))
	}
}
