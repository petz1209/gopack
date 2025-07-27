package logger

import "fmt"

type Logger struct{}

func (l *Logger) Info(msg string) {
	fmt.Println("INFO: " + msg)

}

func (l *Logger) Error(msg string) {
	fmt.Println("ERROR: " + msg)

}
