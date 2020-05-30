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
	switch {
	case err != nil:
		os.Exit(1)
	case string(response) == "y":
		// TODO request to server
	case string(response) == "N" || string(response) == "n":
		break
	default:
		reload_process()
	}
}

func status_process() {
	// TODO request to server
}
