package main

import (
    "fmt"
    "os/signal"
    "os"
    "syscall"
    "github.com/kyazdani42/taskmaster/lib"
)

func load_signals(conf *config.Config) {
	c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGHUP)
    for range(c) {
        new_conf, err := config.GetServerConfig()
        if err != nil {
            fmt.Println(err)
            return
        }
        *conf = new_conf
        fmt.Println(conf)
    }
}

func main() {
    conf, err := config.GetServerConfig()
    if err != nil {
        fmt.Println(err)
        return
    }

    go load_signals(&conf)

    for {}
}
