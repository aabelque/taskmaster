package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	env := parseFlags()
	conf, err := getServerConfig(env.config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go loadSignals(&conf, &env)

	if env.verboseConf {
		conf.Print()
	}

	logger := NewLogger(env.verboseLogs, os.Stdout)
	procs := make(Procs)
	populateProcesses(conf, logger, &procs)
	initProcesses(&procs, logger)

	http.HandleFunc("/reload", unimplemented)
	http.HandleFunc("/status", status(&procs))
	http.HandleFunc("/start", unimplemented)
	http.HandleFunc("/restart", unimplemented)
	http.HandleFunc("/stop", unimplemented)
	http.HandleFunc("/list_progs", listPrograms(&conf))

	http.ListenAndServe(":"+strconv.FormatUint(uint64(conf.Taskmasterd.Port), 10), nil)
}

type Env struct {
	config      string
	verboseConf bool
	verboseLogs bool
}

func parseFlags() Env {
	config := flag.String("config", "", "configuration file path")
	verbose := flag.String("verbose", "", "one of conf, logs or all")
	flag.Parse()

	verboseConf := false
	verboseLogs := false

	switch *verbose {
	case "all":
		verboseConf = true
		verboseLogs = true
	case "conf":
		verboseConf = true
	case "logs":
		verboseLogs = true
	}

	return Env{
		config:      *config,
		verboseConf: verboseConf,
		verboseLogs: verboseLogs,
	}
}

type ProcState struct {
	cmd   Command
	p     *os.Process
	state string
	info  string
	start *time.Time
}

func (p *ProcState) update(proc *os.Process, state string, info string, start *time.Time) {
	p.p = proc
	p.state = state
	p.info = info
	if start != nil {
		p.start = start
	}
}

func initProc(cmd Command) *ProcState {
	return &ProcState{
		cmd:   cmd,
		p:     nil,
		state: "STOPPED",
		info:  "Not started",
		start: nil,
	}
}

type Procs = map[string]*ProcState

func populateProcesses(conf Config, logger Logger, procs *Procs) {
	for cmdName, cmd := range conf.Program {
		if cmd.Instances > 1 {
			var i uint
			for i = 0; i < cmd.Instances; i++ {
				procName := fmt.Sprintf("%s-%d", cmdName, i)
				(*procs)[procName] = initProc(cmd)
			}
		} else {
			(*procs)[cmdName] = initProc(cmd)
		}
	}
}

func initProcesses(procs *Procs, logger Logger) {
	for name, state := range *procs {
		if state.cmd.Startup {
			go runProcAndMonitor(state.cmd, name, &logger, procs)
		}
	}
}

func runProcAndMonitor(cmd Command, name string, l *Logger, procs *Procs) {
	proc, err := cmd.run(name)
	if err != nil {
		(*procs)[name].update(nil, "FATAL", err.err.Error(), nil)
		l.LogActivity(err.logLevel, err.logType, err.err.Error())
	} else {
		startTime := time.Now()
		(*procs)[name].update(proc, "RUNNING", fmt.Sprintf("pid %d, runtime %s", proc.Pid, "0:00:00"), &startTime)
		l.LogActivity("INFO", "spawned", fmt.Sprintf("'%s' with pid %d", name, proc.Pid))
		go cmd.monitor(name, proc, l, procs)
	}
}
