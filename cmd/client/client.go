package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Quit = "quit\n"
)

func main() {
	// set_raw_mode()
	// unset_raw_mode()
	Reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Taskmaster> ")
		text, err := Reader.ReadString('\n')
		fmt.Print(text)
		if err != nil {
			os.Exit(1)
		}
		switch {
		case text == Quit:
			os.Exit(0)
		}
	}
}
