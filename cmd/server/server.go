package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

type Env struct {
	config       string
	verbose_conf bool
	verbose_logs bool
}

func parse_flags() Env {
	config := flag.String("config", "", "configuration file path")
	verbose := flag.String("verbose", "", "one of conf, logs or all")
	flag.Parse()

	verbose_conf := false
	verbose_logs := false

	switch *verbose {
	case "all":
		verbose_conf = true
		verbose_logs = true
	case "conf":
		verbose_conf = true
	case "logs":
		verbose_logs = true
	}

	return Env{
		config:       *config,
		verbose_conf: verbose_conf,
		verbose_logs: verbose_logs,
	}
}

func main() {
	env := parse_flags()
	conf, err := get_server_config(env.config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go load_signals(&conf, &env)

	if env.verbose_conf {
		conf.Print()
	}

	for cmd_name, cmd := range conf.Program {
		if cmd.Startup {
			if cmd.Instances > 1 {
				var i uint
				for i = 0; i < cmd.Instances; i++ {
					go cmd.run(fmt.Sprintf("%s-%d", cmd_name, i))
				}
			} else {
				go cmd.run(cmd_name)
			}
		}
	}

	http.HandleFunc("/reload", unimplemented)
	http.HandleFunc("/status", unimplemented)
	http.HandleFunc("/start", unimplemented)
	http.HandleFunc("/restart", unimplemented)
	http.HandleFunc("/stop", unimplemented)
	http.HandleFunc("/list_progs", program_lister(&conf))

	http.ListenAndServe(":3000", nil)
}

func program_lister(config *Config) func(http.ResponseWriter, *http.Request) {
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
