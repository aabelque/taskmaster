package main

import (
	"encoding/json"
	"fmt"
	"github.com/kyazdani42/taskmaster/pkg/lib"
	"net/http"
	"time"
)

func fmtDuration(d time.Duration) string {
	h := int64(d.Hours())
	m := int64(d.Minutes()) % 60
	s := int64(d.Seconds()) % 60
	return fmt.Sprintf("%d:%02d:%02d", h, m, s)
}

func status(procs *Procs) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		var status []lib.Status
		for name, state := range *procs {
			if state.state == "RUNNING" {
				since := time.Since(*state.start)
				status = append(status, lib.Status{
					Name:   name,
					Status: state.state,
					Info:   fmt.Sprintf("pid %d, runtime %s", state.p.Pid, fmtDuration(since)),
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
