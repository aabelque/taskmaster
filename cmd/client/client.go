package main

import (
	"io"
	"log"
	"os"

	"github.com/chzyer/readline"
)

const (
	help    = "help"
	status  = "status"
	start   = "start"
	restart = "restart"
	reload  = "reload"
	stop    = "stop"
	quit    = "quit"
)

var completer = readline.NewPrefixCompleter(
	readline.PcItem("status"),
	readline.PcItem("start"),
	readline.PcItem("restart"),
	readline.PcItem("reload"),
	readline.PcItem("stop"),
	readline.PcItem("quit"),
)

func main() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:            "taskmaster> ",
		HistoryFile:       ".history_taskmaster",
		InterruptPrompt:   "^C",
		EOFPrompt:         "",
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
		parser_line(line)
	}
}
