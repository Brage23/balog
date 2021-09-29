//go:build rel

package balog

func Panic(a ...interface{}) {
	logPanic(a...)
}

func Error(a ...interface{}) {}

func Warning(a ...interface{}) {}

func Info(a ...interface{}) {}
