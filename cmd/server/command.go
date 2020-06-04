package main

import (
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Command          string
	Instances        uint
	Startup          bool
	Reload           string
	Return_value     int
	Valid_after      uint32
	Kill_after       int
	Closing_signal   string // should be an int ?
	Wait_before_kill int
	Stdout           []string
	Stderr           []string
	Env              []string
	Cwd              string
	Umask            int // maybe a string ?
}

func (c Command) run(name string) {
	args := strings.Split(c.Command, " ")

	cmd := args[0]
	files := make([]*os.File, 3, 3)

	closed, err := os.OpenFile(os.DevNull, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log_activity("CRIT", "oserror", "could not open "+os.DevNull)
		return
	}
	defer closed.Close()

	files[0] = closed
	files[1] = closed
	files[2] = closed

	attr := os.ProcAttr{Dir: c.Cwd, Env: c.Env, Files: files}

	proc, err := os.StartProcess(cmd, args, &attr)
	if err != nil {
		log_activity("INFO", "spawnerr", err.Error())
		return
	}
	log_activity("INFO", "spawned", fmt.Sprintf("'%s' with pid %d", name, proc.Pid))

	ret, err := proc.Wait()
	if err != nil {
		log_activity("INFO", "exited", fmt.Sprintf("%s %s", name, err.Error()))
		return
	}

	var expected string
	time_ok := c.Valid_after <= uint32(ret.SystemTime().Milliseconds())
	if ret.ExitCode() == c.Return_value && time_ok {
		expected = "expected"
	} else {
		expected = "not expected"
	}
	log_activity("INFO", "exited", fmt.Sprintf("%s (exit status %d; %s)", name, ret.ExitCode(), expected))
}
