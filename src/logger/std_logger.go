package logger

import "fmt"

type StdLogger struct {
}

func (stdLogger *StdLogger) PrintLog(level string, message string) {
	logString := fmt.Sprintf("Level: %s, Message: %s", level, message)
	fmt.Println(logString)
}
