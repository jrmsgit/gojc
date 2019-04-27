// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	"github.com/jrmsdev/gojc/internal/logger"
)

var l *logger.Logger

func Init(level string) error {
	var err error
	l, err = logger.New(level)
	return err
}

func Levels() string {
	return logger.Levels
}

func Debug(format string, args ...interface{}) {
	l.Debug(format, args...)
}

func Warnf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}

func Warn(err error) {
	l.Warn(err)
}

func Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}

func Error(err error) {
	l.Error(err)
}

func Printf(format string, args ...interface{}) {
	l.Printf(format, args...)
}

func Print(args ...interface{}) {
	l.Print(args...)
}
