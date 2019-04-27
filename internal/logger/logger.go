// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger // github.com/jrmsdev/gojc/internal/logger

import (
	"errors"
	"io"
	"os"
	"runtime"
	"strings"
)

type Logger struct {
	Debug  func(string, ...interface{})
	Errorf func(string, ...interface{})
	Error  func(error)
	Warnf  func(string, ...interface{})
	Warn   func(error)
	Printf func(string, ...interface{})
	Print  func(...interface{})
}

var shortIdx int

var outfh io.Writer = os.Stdout
var errfh io.Writer = os.Stderr

var Levels string = "debug, warn, error or quiet"

func New(level string) (*Logger, error) {
	l := new(Logger)
	l.Debug = quietf
	l.Warnf = quietf
	l.Warn = quieterr
	// default level: error
	l.Errorf = errorf
	l.Error = perror
	l.Printf = printf
	l.Print = print
	if level == "debug" {
		// debug also enables warn
		l.Debug = debug
		l.Warnf = warnf
		l.Warn = warn
	} else if level == "warn" {
		l.Warnf = warnf
		l.Warn = warn
	} else if level == "quiet" {
		l.Printf = quietf
		l.Print = quiet
	} else if level != "error" {
		return nil, errors.New("invalid logger level " + level)
	}
	shortIdx = getShortIdx()
	return l, nil
}

func getShortIdx() int {
	_, fn, _, ok := runtime.Caller(0)
	if ok {
		return strings.Index(fn, "github.com")
	}
	return 0
}

func Testing(w io.Writer) {
	outfh = w
	errfh = w
}
