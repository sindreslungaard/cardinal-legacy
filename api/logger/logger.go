package logger

import (
	"fmt"
	"time"
)

func Debug(msg string, args ...interface{}) {
	log(format(msg))
}

func Info(msg string, args ...interface{}) {
	log(format(msg))
}

func Warn(msg string, args ...interface{}) {
	log(format(msg))
}

func Fatal(msg string, args ...interface{}) {
	log(format(msg))
}

func format(msg string, args ...interface{}) string {
	return fmt.Sprintf(msg, args...)
}

func log(msg string) {
	t := time.Now()

	prefix := t.Format("15:04:05")

	println(fmt.Sprintf("%s %s", prefix, msg))
}
