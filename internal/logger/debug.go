// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

import (
	"fmt"
	"os"
	"runtime"
)

func debug(format string, args ...interface{}) {
	tag := "D: "
	_, fn, ln, ok := runtime.Caller(2)
	if ok {
		tag = fmt.Sprintf("%s:%d: ", fn[shortIdx:], ln)
	}
	fmt.Fprintf(os.Stderr, tag+format+"\n", args...)
}
