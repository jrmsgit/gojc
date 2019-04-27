// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

import (
	"fmt"
)

func printf(format string, args ...interface{}) {
	fmt.Fprintf(outfh, format+"\n", args...)
}

func print(args ...interface{}) {
	fmt.Fprint(outfh, args...)
	fmt.Fprint(outfh, "\n")
}
