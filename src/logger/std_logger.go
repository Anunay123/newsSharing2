package logger

import "fmt"

type StdLogger struct {
}

func (stdLogger *StdLogger) PrintLog(level string, message string) {
	fmt.Printf("Level: %s, Message: %s", level, message)
}
