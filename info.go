//go:build !error && !warning && !rel
// +build !error,!warning,!rel

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

func Info(a ...interface{}) {
	logInfo(a...)
}
