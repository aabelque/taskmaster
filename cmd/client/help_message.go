package main

import (
	"fmt"
	"strings"
)

func restartMessage() {
	var message strings.Builder
	message.WriteString("\x1b[31mError\x1b[0m: relaunch requires a process name\n")
	message.WriteString("relaunch <name>          Relaunch a process\n")
	message.WriteString("relaunch <name> <name>   Relaunch multiple processes\n")
	message.WriteString("relaunch all             Relaunch all processes")
	fmt.Println(message.String())
}

func startMessage() {
	var message strings.Builder
	message.WriteString("\x1b[31mError\x1b[0m: launch requires a process name\n")
	message.WriteString("launch <name>          Launch a process\n")
	message.WriteString("launch <name> <name>   Launch multiple processes\n")
	message.WriteString("launch all             Launch all processes")
	fmt.Println(message.String())
}

func stopMessage() {
	var message strings.Builder
	message.WriteString("\x1b[31mError\x1b[0m: stop requires a process name\n")
	message.WriteString("stop <name>          Stop a process\n")
	message.WriteString("stop <name> <name>   Stop multiple processes\n")
	message.WriteString("stop all             Stop all processes")
	fmt.Println(message.String())
}

func helpMessage() {
	var message strings.Builder
	message.WriteString("\ndefault commands:\n")
	message.WriteString("=================\n")
	message.WriteString("help     status  launch\n")
	message.WriteString("relaunch reload  stop\n")
	message.WriteString("quit     list\n")
	fmt.Println(message.String())
}
