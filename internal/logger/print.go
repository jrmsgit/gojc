// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

import (
	"fmt"
)

func printf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func print(args ...interface{}) {
	fmt.Print(args...)
	fmt.Print("\n")
}
