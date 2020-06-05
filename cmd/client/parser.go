package main

import (
	"os"
	"strings"
)

func check_command(args []string, c Config) {
	command := strings.TrimSpace(args[0])
	switch command {
	case quit:
		os.Exit(0)
	case help:
		help_message()
	case reload:
		reload_process(c)
	case status:
		status_process(c)
	case start:
		start_message()
	case restart:
		restart_message()
	case stop:
		stop_message()
	default:
		break
	}
}

func parser_line(line string, c Config) {
	args := strings.Split(strings.TrimSpace(line), " ")
	if len(args) == 1 {
		check_command(args, c)
	} else {
		command, process := args[0], args[1:]
		choose_command(command, process, c)
	}
}
