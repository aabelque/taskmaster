package main

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	log func(string, ...interface{})
}

func NewLogger(log bool, file *os.File) Logger {
	if log {
		return Logger{
			log: func(s string, v ...interface{}) {
				fmt.Fprintf(file, s, v...)
			},
		}
	} else {
		return Logger{
			log: func(s string, v ...interface{}) {},
		}
	}
}

func (l *Logger) LogActivity(level string, val string, content string) {
	t := time.Now().Format("2006-01-02 15:04:05.000")
	l.log("%s %s \x1b[1m%s\x1b[0m: %s\n", t, level, val, content)
}
