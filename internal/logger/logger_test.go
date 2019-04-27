// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"
)

var buf *bytes.Buffer

func TestMain(m *testing.M) {
	buf = new(bytes.Buffer)
	Testing(buf)
	callerSkip = 1
	rc := m.Run()
	buf.Reset()
	os.Exit(rc)
}

func check(t *testing.T, expect string) {
	t.Helper()
	got := strings.TrimSpace(buf.String())
	buf.Reset()
	if !strings.HasPrefix(got, expect) {
		t.Errorf("got='%s' expect='%s'", got, expect)
	}
}

func TestDebug(t *testing.T) {
	l, err := New("debug")
	if err != nil {
		t.Fatal(err)
	}
	l.Debug("tes%sing", "t")
	check(t, "github.com/jrmsdev/gojc/internal/logger/logger_test.go:")
}

func TestWarn(t *testing.T) {
	l, err := New("warn")
	if err != nil {
		t.Fatal(err)
	}
	l.Warn(errors.New("testing"))
	check(t, "W: testing")
	l.Warnf("tes%sing", "t")
	check(t, "W: testing")
}

func TestError(t *testing.T) {
	l, err := New("error")
	if err != nil {
		t.Fatal(err)
	}
	l.Error(errors.New("testing"))
	check(t, "E: testing")
	l.Errorf("tes%sing", "t")
	check(t, "E: testing")
}

func TestPrint(t *testing.T) {
	l, err := New("error")
	if err != nil {
		t.Fatal(err)
	}
	l.Print(errors.New("testing"))
	check(t, "testing")
	l.Printf("tes%sing", "t")
	check(t, "testing")
}

func TestQuiet(t *testing.T) {
	l, err := New("quiet")
	if err != nil {
		t.Fatal(err)
	}
	l.Debug("tes%sing", "t")
	check(t, "")
	l.Warnf("tes%sing", "t")
	check(t, "")
	l.Errorf("tes%sing", "t")
	check(t, "E: testing") // errors are printed even in quiet mode
	l.Print(errors.New("testing"))
	check(t, "")
	l.Printf("tes%sing", "t")
	check(t, "")
}

func TestLevel(t *testing.T) {
	_, err := New("fake")
	if err == nil {
		t.Fatal("log level 'fake' should have failed")
	}
}
