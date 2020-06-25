package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

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

	for cmdName, cmd := range conf.Program {
		if cmd.Startup {
			if cmd.Instances > 1 {
				var i uint
				for i = 0; i < cmd.Instances; i++ {
					go cmd.run(fmt.Sprintf("%s-%d", cmdName, i), &logger)
				}
			} else {
				go cmd.run(cmdName, &logger)
			}
		}
	}

	http.HandleFunc("/reload", unimplemented)
	http.HandleFunc("/status", unimplemented)
	http.HandleFunc("/start", unimplemented)
	http.HandleFunc("/restart", unimplemented)
	http.HandleFunc("/stop", unimplemented)
	http.HandleFunc("/list_progs", programLister(&conf))

	http.ListenAndServe(":"+strconv.FormatUint(uint64(conf.Taskmasterd.Port), 10), nil)
}

func programLister(config *Config) func(http.ResponseWriter, *http.Request) {
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
