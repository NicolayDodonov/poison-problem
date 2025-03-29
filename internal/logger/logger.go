package logger

import (
	"log"
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
	l := &Logger{path, convert(level)}
	l.write("+====================+")
	l.Info("Logger Init")
	return l
}

func (l *Logger) Error(str string) {
	if l.Level <= Error {
		t := time.Now()

		l.write("[ERR] " + t.Format("02-01-2006 15:04:05") + "; " + str + ";")
	}
}

func (l *Logger) Debug(str string) {
	if l.Level <= Debug {
		t := time.Now()

		l.write("[DEB] " + t.Format("02-01-2006 15:04:05") + "; " + str + ";")
	}
}

func (l *Logger) Info(str string) {
	if l.Level <= Info {
		t := time.Now()

		l.write("[INF] " + t.Format("02-01-2006 15:04:05") + "; " + str + ";")
	}
}

// it is not good realisation of write to log file, but
func (l *Logger) write(msg string) {
	f, err := os.OpenFile(l.Path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(msg + "\n"); err != nil {
		//it not stable!!!!!!!!!!!!!!!
		log.Printf("Logger error " + err.Error())
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
