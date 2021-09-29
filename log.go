package balog

import (
	"fmt"
	"time"
)

func currtime() string {
	return fmt.Sprintf("[%d-%d-%d %d:%d:%d.%d]", time.Now().Year(), time.Now().Month(), time.Now().Day(),
		time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond())
}

func log(title string, a ...interface{}) {
	fmt.Print(currtime())
	fmt.Print(" ", title, " - ")
	fmt.Println(a...)
}

func logPanic(a ...interface{}) {
	log("Panic", a...)
}

func logError(a ...interface{}) {
	log("Error", a...)
}

func logWarning(a ...interface{}) {
	log("Warning", a...)
}

func logInfo(a ...interface{}) {
	log("Info", a...)
}
