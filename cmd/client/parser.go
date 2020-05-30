package main

import (
	"fmt"
	"os"
	"strings"
)

func check_command(command string) {
	switch command {
	case quit:
		os.Exit(0)
	case help:
		help_message()
	case reload:
		reload_process()
	case status:
		// TODO status_process()
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

func parser_line(line string) {
	args := strings.Split(strings.TrimSpace(line), " ")
	if len(args) == 1 {
		check_command(strings.TrimSpace(args[0]))
	} else {
		command, process := args[0], args[1:]
		//TODO request to server
		fmt.Println(command, process)
	}
}
