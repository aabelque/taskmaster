package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/chzyer/readline"
)

type Config struct {
	Serverurl   string
	Port        uint16
	Username    string
	Password    string
	Prompt      string
	HistoryFile string
}

var cli bool = false
var commands string = "help status start restart reload stop quit list"

const (
	help    = "help"
	status  = "status"
	start   = "start"
	restart = "restart"
	reload  = "reload"
	stop    = "stop"
	quit    = "quit"
	list    = "list"
)

var completer = readline.NewPrefixCompleter(
	readline.PcItem(status),
	readline.PcItem(start),
	readline.PcItem(restart),
	readline.PcItem(reload),
	readline.PcItem(stop),
	readline.PcItem(quit),
	readline.PcItem(list),
)

func prompt(c Config) {
	l, err := readline.NewEx(&readline.Config{
		Prompt:            c.Prompt,
		HistoryFile:       c.HistoryFile,
		InterruptPrompt:   "^C",
		EOFPrompt:         "quit",
		AutoComplete:      completer,
		HistorySearchFold: true,
	})
	if err != nil {
		log.Fatalln(err)
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
		parserLine(line, c)
	}
}

func commandLine(c Config) {
	if len(os.Args) == 2 {
		checkCommand(os.Args[1:], c)
	} else {
		command, process := os.Args[1], os.Args[2:]
		if ret, _ := regexp.Match(command, []byte(commands)); !ret {
			cerr := command
			for i := 0; i < len(process); i++ {
				cerr += " " + process[i]
			}
			fmt.Println("*** Unknown syntax: " + cerr)
		} else if command == reload {
			reloadProcess(command, process, c)
		} else {
			request(command, process, c)
		}
	}
}

func main() {
	var conf Config
	getClientConfig(&conf)
	if len(os.Args) > 1 {
		cli = true
		commandLine(conf)
	} else {
		listProgs("/list_progs", conf)
		prompt(conf)
	}
}
