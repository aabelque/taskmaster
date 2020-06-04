package main

import (
	"bufio"
	"fmt"
	"os"
)

func reload_process() {
	if !cli {
		fmt.Println("Really restart the remote taskmasterd process y/N?")
		buf := bufio.NewReader(os.Stdin)
		response, err := buf.ReadByte()
		if err != nil {
			return
		}
		switch string(response) {
		case "y":
			// TODO request to server
		case "n":
			break
		case "N":
			break
		default:
			reload_process()
		}
	} else {
		// TODO request to server
	}
}

func status_process() {
	// TODO request to server
}

func start_process(process []string) {
	// TODO request to server
}

func restart_process(process []string) {
	// TODO request to server
}

func stop_process(process []string) {
	// TODO request to server
}

func choose_command(command string, process []string) {
	switch command {
	case start:
		start_process(process)
	case restart:
		restart_process(process)
	case stop:
		stop_process(process)
	case status:
		status_process()
	default:
		break
	}
}
