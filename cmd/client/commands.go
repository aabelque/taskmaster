package main

import (
	"bufio"
	"fmt"
	// "log"
	// "net/http"
	// "net/http/httputil"
	"os"
)

func reload_process(c Config) {
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
			reload_process(c)
		}
	} else {
		// TODO request to server
	}
}

func status_process(c Config) {
	// TODO request to server
}

func start_process(process []string, c Config) {
	// TODO request to server
}

func restart_process(process []string, c Config) {
	// TODO request to server
}

func stop_process(process []string, c Config) {
	// TODO request to server
}

func choose_command(command string, process []string, c Config) {
	switch command {
	case start:
		start_process(process, conf)
	case restart:
		restart_process(process, conf)
	case stop:
		stop_process(process, conf)
	case status:
		status_process(conf)
	default:
		break
	}
}

func request(command string, process []string, c Config) {
}
