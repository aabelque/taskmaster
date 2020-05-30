package main

import (
	"fmt"
	"os"
	"strings"
)

func check_command(command string) {
	switch {
	case command == quit:
		os.Exit(0)
	case command == help:
		help_message()
	case command == reload:
		reload_process()
	case command == status:
		// TODO status_process()
	case command == start:
		start_message()
	case command == restart:
		restart_message()
	case command == stop:
		stop_message()
	default:
		break
	}
}

func parser_line(line string) {
	args := strings.Split(line, " ")
	if len(args) == 1 {
		check_command(strings.TrimSpace(args[0]))
	} else {
		command, process := args[0], args[1:]
		if len(process) == 1 {
			x := strings.TrimSpace(process[0])
			if len(x) == 0 {
				check_command(command)
			}
			//TODO request to server
		} else {
			//TODO request to server
			fmt.Println(command, process)
		}
	}
}
