package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

//TODO refactor checkCommand and parserLine

func checkCommand(args []string, c Config) {
	command, process := strings.TrimSpace(args[0]), args[1:]
	switch command {
	case quit:
		os.Exit(0)
	case help:
		helpMessage()
	case reload:
		reloadProcess(command, process, c)
	case status:
		request(command, process, c)
	case start:
		startMessage()
	case restart:
		restartMessage()
	case stop:
		stopMessage()
	case list:
		listProgs("/list_progs", c)
	default:
		fmt.Println("*** Unknown syntax: " + command)
	}
}

func parserLine(line string, c Config) {
	args := strings.Split(strings.TrimSpace(line), " ")
	if len(args) == 1 {
		checkCommand(args, c)
	} else {
		command, process := args[0], args[1:]
		if ret, _ := regexp.Match(command, []byte(commands)); !ret {
			cerr := command
			for i := 0; i < len(process); i++ {
				cerr += " " + process[i]
			}
			fmt.Println("*** Unknown syntax: " + cerr)
		}
		if command == reload {
			reloadProcess(command, process, c)
		} else {
			request(command, process, c)
		}
	}
}

func parserRequest(req []byte) string {
	size := len(req)
	var str string
	for i := 0; i < size; i++ {
		switch string(req[i]) {
		case "[":
			continue
		case ",":
			str += " | "
		case "]":
			continue
		case "\"":
			continue
		default:
			str += string(req[i])
		}
	}
	return str
}
