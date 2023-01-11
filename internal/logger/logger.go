package logger

import (
	"fmt"
	"log"
)

func NewLogger() Logger {
	return &logger{
		l: log.Default(),
	}
}

type Logger interface {
	Errorf(format string, v ...any)
	Info(format string, v ...any)
}

type logger struct {
	l *log.Logger
}

func (l *logger) Errorf(format string, v ...any) {
	if len(v) == 0 {
		l.l.Printf("[Error]: %s\n", format)
		return
	}
	msg := fmt.Sprintln(format, v)
	l.l.Printf("[Error]: %s\n", msg)
}

func (l *logger) Info(format string, v ...any) {
	if len(v) == 0 {
		l.l.Printf("[Info]: %s\n", format)
		return
	}
	msg := fmt.Sprintln(format, v)
	l.l.Printf("[Info]: %s\n", msg)
}
