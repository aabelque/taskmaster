package main

import (
	"fmt"
	"strings"
)

func restart_message() {
	var message strings.Builder
	message.WriteString("\x1b[31mError\x1b[0m: restart requires a process name\n")
	message.WriteString("restart <name>          Restart a process\n")
	message.WriteString("restart <name> <name>   Restart multiple processes\n")
	message.WriteString("restart all             Restart all processes")
	fmt.Println(message.String())
}

func start_message() {
	var message strings.Builder
	message.WriteString("\x1b[31mError\x1b[0m: start requires a process name\n")
	message.WriteString("start <name>          Start a process\n")
	message.WriteString("start <name> <name>   Start multiple processes\n")
	message.WriteString("start all             Start all processes")
	fmt.Println(message.String())
}

func stop_message() {
	var message strings.Builder
	message.WriteString("\x1b[31mError\x1b[0m: stop requires a process name\n")
	message.WriteString("stop <name>          Stop a process\n")
	message.WriteString("stop <name> <name>   Stop multiple processes\n")
	message.WriteString("stop all             Stop all processes")
	fmt.Println(message.String())
}

func help_message() {
	var message strings.Builder
	message.WriteString("\ndefault commands:\n")
	message.WriteString("=================\n")
	message.WriteString("help     status  launch\n")
	message.WriteString("relaunch reload  stop\n")
	message.WriteString("quit\n")
	fmt.Println(message.String())
}
