package main

import (
	"bufio"
	"fmt"
	"os"
)

func reload_process() {
	fmt.Println("Really restart the remote supervisord process y/N?")
	buf := bufio.NewReader(os.Stdin)
	response, err := buf.ReadByte()
	if err != nil {
		os.Exit(1)
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
}

func status_process() {
	// TODO request to server
}
