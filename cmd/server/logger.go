package main

import (
	"fmt"
	"time"
)

func log_activity(level string, val string, content string) {
	t := time.Now().Format("2006-01-02 15:04:05.000")
	fmt.Printf("%s %s \x1b[1m%s\x1b[0m: %s\n", t, level, val, content)
}
