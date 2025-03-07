package logger

import (
	"os"
	"time"
)

type Level uint8

type Logger struct {
	Path  string
	Level Level
}

const (
	Debug Level = iota
	Info
	Error
	Off = 10
)

func New(path string, level Level) *Logger {
	return &Logger{path, level}
}

func (l *Logger) Error(str string) {
	if l.Level <= Error {
		t := time.Now()

		msg := "[ERR] " + str + " Time:" + t.Format("02-01-2006 15:04:05")
		_ = os.WriteFile(l.Path, []byte(msg), 0666)
	}
}

func (l *Logger) Debug(str string) {
	if l.Level <= Debug {
		t := time.Now()

		msg := "[DEB] " + str + " Time:" + t.Format("02-01-2006 15:04:05")
		_ = os.WriteFile(l.Path, []byte(msg), 0666)
	}
}

func (l *Logger) Info(str string) {
	if l.Level <= Info {
		t := time.Now()

		msg := "[INF] " + str + " Time:" + t.Format("02-01-2006 15:04:05")
		_ = os.WriteFile(l.Path, []byte(msg), 0666)
	}
}
