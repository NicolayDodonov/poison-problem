package logger

import (
	"os"
	"strings"
	"time"
)

type level uint8

type Logger struct {
	Path  string
	Level level
}

const (
	Debug level = iota
	Info
	Error
	Off = 10
)

func New(path, level string) *Logger {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		//if file log is not found - create!
		_, _ = os.Create(path)
	}
	return &Logger{path, convert(level)}
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

func convert(s string) level {
	switch strings.ToLower(s) {
	case "debug":
		return Debug
	case "info":
		return Info
	case "error":
		return Error
	case "off":
		return Off
	default:
		return Off
	}
}
