package main

import (
	"fmt"
	"github.com/kyazdani42/taskmaster/pkg/lib"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Command struct {
	Command          string
	Instances        uint
	Startup          bool
	Reload           string
	Return_value     int
	Valid_after      int
	Kill_after       int
	Closing_signal   string // should be an int ?
	Wait_before_kill int
	Stdout           []string
	Stderr           []string
	Env              []string
	Cwd              string
	Umask            int // maybe a string ?
}

type Config = map[string]map[string]Command

func get_server_config() (Config, error) {
	var conf Config
	if err := config.GetConfig(&conf, "config.toml"); err != nil {
		return conf, err
	}
	return conf, nil
}

func load_signals(conf *Config) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)
	for range c {
		new_conf, err := get_server_config()
		if err != nil {
			fmt.Println(err)
			return
		}
		*conf = new_conf
		fmt.Println(conf)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	conf, err := get_server_config()
	if err != nil {
		fmt.Println(err)
		return
	}

	go load_signals(&conf)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
