package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Command        string
	Instances      uint
	Startup        bool
	Reload         string
	ReturnValue    int
	ValidAfter     uint32
	KillAfter      int
	ClosingSignal  string // should be an int ?
	WaitBeforeKill int
	Stdout         []string
	Stderr         []string
	Env            []string
	Cwd            string
	Umask          int // maybe a string ?
}

func getFile(v []string) (*os.File, error) {
	var fd *os.File
	var err error

	switch v[0] {
	case "normal":
		fd = os.Stdout
		break
	case "close":
		fd, err = os.OpenFile(os.DevNull, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		break
	case "redirect":
		fd, err = os.OpenFile(v[1], os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		break
	default:
		return nil, errors.New("Wrong config option")
	}

	if err != nil {
		return nil, err
	} else {
		return fd, nil
	}
}

func getFiles(stdout []string, stderr []string) (*os.File, *os.File, error) {
	outFd, err := getFile(stdout)
	if err != nil {
		return nil, nil, err
	}

	errFd, err := getFile(stderr)
	if err != nil {
		return nil, nil, err
	}

	return outFd, errFd, nil
}

func (c Command) run(name string, l *Logger) (*os.Process, error) {
	args := strings.Split(c.Command, " ")

	cmd := args[0]
	files := make([]*os.File, 3, 3)

	fdOut, fdErr, err := getFiles(c.Stdout, c.Stderr)
	if err != nil {
		l.LogActivity("CRIT", "fileerror", err.Error())
		return nil, nil
	}
	defer fdOut.Close()
	defer fdErr.Close()

	files[1] = fdOut
	files[2] = fdErr

	attr := os.ProcAttr{Dir: c.Cwd, Env: c.Env, Files: files}

	return os.StartProcess(cmd, args, &attr)
}

func (c Command) monitor(name string, proc *os.Process, l *Logger) {
	ret, err := proc.Wait()
	if err != nil {
		l.LogActivity("INFO", "exited", fmt.Sprintf("%s %s", name, err.Error()))
		return
	}

	var expected string
	timeOk := c.ValidAfter <= uint32(ret.SystemTime().Milliseconds())
	if ret.ExitCode() == c.ReturnValue && timeOk {
		expected = "expected"
	} else {
		expected = "not expected"
	}
	l.LogActivity("INFO", "exited", fmt.Sprintf("%s (exit status %d; %s)", name, ret.ExitCode(), expected))
}
