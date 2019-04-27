// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

import (
	"fmt"
)

func warnf(format string, args ...interface{}) {
	fmt.Fprintf(errfh, "W: "+format+"\n", args...)
}

func warn(err error) {
	warnf("%s", err)
}
