//go:build rel

package balog

import (
	"fmt"
	"os"
	"runtime"
)

func Error(a ...interface{}) {}

func Warning(a ...interface{}) {}

func Info(a ...interface{}) {}

func RelAssert(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		assert_position := fmt.Sprintf("[%s:%d]", file, line)
		log("RelAssert", assert_position)
		os.Exit(1)
	}
}

func Assert(err error) {}
