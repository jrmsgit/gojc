// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

import (
	"fmt"
	"runtime"
)

var callerSkip int = 2

func debug(format string, args ...interface{}) {
	tag := "D: "
	_, fn, ln, ok := runtime.Caller(callerSkip)
	if ok {
		tag = fmt.Sprintf("%s:%d: ", fn[shortIdx:], ln)
	}
	fmt.Fprintf(errfh, tag+format+"\n", args...)
}
