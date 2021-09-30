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

func RelAssert(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		assert_position := fmt.Sprintf("[%s:%d]", file, line)
		log("RelAssert", assert_position)
		os.Exit(1)
	}
}

func Assert(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		assert_position := fmt.Sprintf("[%s:%d]", file, line)
		log("Assert", assert_position)
		os.Exit(1)
	}
}
