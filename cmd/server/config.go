package main

import (
	"fmt"
	"github.com/kyazdani42/taskmaster/pkg/lib"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	Taskmasterd Options
	Program     map[string]Command
}

type Options struct {
	Port      uint32
	Directory string
	Loglevel  string
	Logfile   string
}

func getServerConfig(configFile string) (Config, error) {
	var conf Config
	if err := config.GetConfig(&conf, "config.toml", configFile); err != nil {
		return Config{}, err
	}

	return conf, nil
}

func loadSignals(conf *Config, env *Env) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)
	for range c {
		newConf, err := getServerConfig(env.config)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		*conf = newConf
		if env.verboseConf {
			conf.Print()
		}
	}
}

func (c *Config) Print() {
	fmt.Println("\x1b[1mconfig: {\x1b[0m")
	fmt.Println("  directory:", c.Taskmasterd.Directory)
	fmt.Println("  log file:", c.Taskmasterd.Logfile)
	fmt.Println("  log level:", c.Taskmasterd.Loglevel)
	fmt.Println("  port:", c.Taskmasterd.Port)
	fmt.Println("\x1b[1m}\x1b[0m")

	fmt.Println()

	fmt.Println("\x1b[1mprograms: {\x1b[0m")
	for progname, cmd := range c.Program {
		fmt.Println("  \x1b[1m" + progname + ": {\x1b[0m")
		fmt.Println("    cmd:", cmd.Command)
		fmt.Println("    cwd:", cmd.Cwd)
		fmt.Println("    instances:", cmd.Instances)
		fmt.Println("    startup:", cmd.Startup)
		fmt.Println("    reload:", cmd.Reload)
		fmt.Println("    expected return:", cmd.ReturnValue)
		fmt.Println("    valid timeout:", cmd.ValidAfter)
		fmt.Println("    kill after timeout:", cmd.KillAfter)
		fmt.Println("    closing signal:", cmd.ClosingSignal)
		fmt.Println("    wait before kill:", cmd.WaitBeforeKill)
		fmt.Println("    stdout:", cmd.Stdout)
		fmt.Println("    stderr:", cmd.Stderr)
		fmt.Println("    env:", cmd.Env)
		fmt.Println("    umask:", cmd.Umask)
		fmt.Println("  \x1b[1m}\x1b[0m")
	}
	fmt.Println("\x1b[1m}\x1b[0m")
}
