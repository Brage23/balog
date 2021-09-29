//go:build warning

package balog

func Panic(a ...interface{}) {
	logPanic(a...)
}

func Error(a ...interface{}) {
	logError(a...)
}

func Warning(a ...interface{}) {
	logWarning(a...)
}

func Info(a ...interface{}) {}
