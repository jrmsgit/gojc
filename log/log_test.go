// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/jrmsdev/gojc/internal/logger"
)

var buf *bytes.Buffer

func TestMain(m *testing.M) {
	buf = new(bytes.Buffer)
	logger.Testing(buf)
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
	err := Init("debug")
	if err != nil {
		t.Fatal(err)
	}
	Debug("tes%sing", "t")
	check(t, "github.com/jrmsdev/gojc/log/log_test.go:")
}

func TestWarn(t *testing.T) {
	err := Init("warn")
	if err != nil {
		t.Fatal(err)
	}
	Warn(errors.New("testing"))
	check(t, "W: testing")
	Warnf("tes%sing", "t")
	check(t, "W: testing")
}

func TestError(t *testing.T) {
	err := Init("error")
	if err != nil {
		t.Fatal(err)
	}
	Error(errors.New("testing"))
	check(t, "E: testing")
	Errorf("tes%sing", "t")
	check(t, "E: testing")
}

func TestPrint(t *testing.T) {
	err := Init("error")
	if err != nil {
		t.Fatal(err)
	}
	Print(errors.New("testing"))
	check(t, "testing")
	Printf("tes%sing", "t")
	check(t, "testing")
}

func TestQuiet(t *testing.T) {
	err := Init("quiet")
	if err != nil {
		t.Fatal(err)
	}
	Debug("tes%sing", "t")
	check(t, "")
	Warnf("tes%sing", "t")
	check(t, "")
	Errorf("tes%sing", "t")
	check(t, "E: testing") // errors are printed even in quiet mode
	Print(errors.New("testing"))
	check(t, "")
	Printf("tes%sing", "t")
	check(t, "")
}

func TestLevels(t *testing.T) {
	lvls := Levels()
	if lvls != "debug, warn, error or quiet" {
		t.Errorf("log levels description: %s", lvls)
	}
}
