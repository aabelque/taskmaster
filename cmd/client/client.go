package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/chzyer/readline"
)

const (
	help     = "help"
	status   = "status"
	launch   = "launch"
	relaunch = "relaunch"
	reload   = "reload"
	stop     = "stop"
	quit     = "quit"
)

var completer = readline.NewPrefixCompleter(
	readline.PcItem("status"),
	readline.PcItem("launch"),
	readline.PcItem("relaunch"),
	readline.PcItem("reload"),
	readline.PcItem("stop"),
	readline.PcItem("quit"),
)

func help_message() {
	var message strings.Builder
	message.WriteString("default commands:\n")
	message.WriteString("=================\n")
	message.WriteString("help\tstatus\tlaunch\n")
	message.WriteString("relaunch\treload\tstop\n")
	message.WriteString("quit\n")

	fmt.Println(message.String())
}

func main() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:            "taskmaster> ",
		HistoryFile:       ".history_taskmaster",
		InterruptPrompt:   "^C",
		EOFPrompt:         "quit",
		AutoComplete:      completer,
		HistorySearchFold: true,
	})
	if err != nil {
		os.Exit(1)
	}
	defer l.Close()

	log.SetOutput(l.Stderr())
	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}
		switch {
		case line == help:
			// request.go
			help_message()
		case line == status:
			// request.go
			log.Println("Status done!") // Ideas to inform user
		case line == launch:
			// request.go
			log.Println("Launch done!") // Ideas to inform user
		case line == relaunch:
			// request.go
			log.Println("Relaunch done!") // Ideas to inform user
		case line == reload:
			// request.go
			log.Println("Reload done!") // Ideas to inform user
		case line == stop:
			// request.go
			log.Println("Stop done!") // Ideas to inform user
		case line == quit:
			os.Exit(0)
		}
	}
}
