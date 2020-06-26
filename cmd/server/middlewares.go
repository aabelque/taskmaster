package main

import (
	"encoding/json"
	"fmt"
	"github.com/kyazdani42/taskmaster/pkg/lib"
	"net/http"
)

func status(procs *map[string]ProcState) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		var status []lib.Status
		for name, state := range *procs {
			if state.info == "RUNNING" {
				status = append(status, lib.Status{
					Name:   name,
					Status: state.state,
					// TODO elapsed time
					Info: fmt.Sprintf("pid %d, runtime %s", state.p.Pid, ""),
				})
			} else {
				status = append(status, lib.Status{
					Name:   name,
					Status: state.state,
					Info:   state.info,
				})
			}
		}
		val, err := json.Marshal(status)
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
