package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
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
	procs := make(map[string]ProcState)
	initProcesses(conf, logger, procs)

	http.HandleFunc("/reload", unimplemented)
	http.HandleFunc("/status", unimplemented)
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
	p     *os.Process
	state string
	info  string
}

func (p ProcState) update(proc *os.Process, state string, info string) {
	p.p = proc
	p.state = state
	p.info = info
}

func (p ProcState) init() {
	p.p = nil
	p.state = "STOPPED"
	p.info = "Not started"
}

func initProcesses(conf Config, logger Logger, procs map[string]ProcState) {
	for cmdName, cmd := range conf.Program {
		if cmd.Instances > 1 {
			var i uint
			for i = 0; i < cmd.Instances; i++ {
				procName := fmt.Sprintf("%s-%d", cmdName, i)
				procs[procName] = ProcState{}
				procs[procName].init()
				if cmd.Startup {
					go runProcAndMonitor(cmd, procName, &logger, &procs)
				}
			}
		} else {
			procs[cmdName] = ProcState{}
			procs[cmdName].init()
			if cmd.Startup {
				go runProcAndMonitor(cmd, cmdName, &logger, &procs)
			}
		}
	}
}

func runProcAndMonitor(cmd Command, name string, l *Logger, procs *map[string]ProcState) {
	proc, err := cmd.run(name)
	if err != nil {
		(*procs)[name].update(nil, "FATAL", err.err.Error())
		l.LogActivity(err.logLevel, err.logType, err.err.Error())
	} else {
		(*procs)[name].update(proc, "RUNNING", fmt.Sprintf("pid %d, runtime %s", proc.Pid, "0:00:00"))
		l.LogActivity("INFO", "spawned", fmt.Sprintf("'%s' with pid %d", name, proc.Pid))
		go cmd.monitor(name, proc, l, procs)
	}
}

func listPrograms(config *Config) func(http.ResponseWriter, *http.Request) {
	keys := make([]string, len(config.Program))
	i := 0
	for k := range config.Program {
		keys[i] = k
		i++
	}

	return func(res http.ResponseWriter, req *http.Request) {
		val, err := json.Marshal(keys)
		if err != nil {
			res.WriteHeader(500)
			return
		}

		res.Write(val)
	}
}

func unimplemented(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Unimplemented !"))
}
