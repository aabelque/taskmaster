package main

import (
	"io"
	"log"
	"os"

	"github.com/chzyer/readline"
)

type Config struct {
	Serverurl    string
	Port         uint16
	Username     string
	Password     string
	Prompt       string
	History_file string
}

var cli bool = false

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
	readline.PcItem(status),
	readline.PcItem(start),
	readline.PcItem(restart),
	readline.PcItem(reload),
	readline.PcItem(stop),
	readline.PcItem(quit),
)

func prompt(c Config) {
	l, err := readline.NewEx(&readline.Config{
		Prompt:            c.Prompt,
		HistoryFile:       c.History_file,
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
		parser_line(line, c)
	}
}

func command_line(c Config) {
	if len(os.Args) == 2 {
		check_command(os.Args[1:], c)
	} else {
		command, process := os.Args[1], os.Args[2:]
		choose_command(command, process, c)
	}
}

func main() {
	var conf Config
	if err := get_client_config(&conf); err != nil {
		os.Exit(1)
	}
	if len(os.Args) > 1 {
		cli = true
		command_line(conf)
	} else {
		prompt(conf)
	}
}
