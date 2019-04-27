// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

import (
	"fmt"
	"os"
)

func errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "E: "+format+"\n", args...)
}

func perror(err error) {
	errorf("%s", err)
}
