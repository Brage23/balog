//go:build error

package balog

func Panic(a ...interface{}) {
	logPanic(a...)
}

func Error(a ...interface{}) {
	logError(a...)
}

func Warning(a ...interface{}) {}

func Info(a ...interface{}) {}
