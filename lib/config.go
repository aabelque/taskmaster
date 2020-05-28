package config

import (
	"github.com/BurntSushi/toml"
	"errors"
	"io/ioutil"
	"os"
	"path"
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

type Config = map[string] map[string] Command

func get_config_folder() (string, error) {
	xdg_config_home, cok := os.LookupEnv("XDG_CONFIG_HOME")
	home, hok := os.LookupEnv("HOME")
	if !cok && !hok {
		return "", errors.New("Could not locate configuration folder")
	} else if !cok {
		return path.Join(home, ".config/taskmaster"), nil
	} else {
		return path.Join(xdg_config_home, "taskmaster"), nil
	}
}

func get_config(conf interface{}, filename string) (error) {
	config_folder, err := get_config_folder()
	if err != nil {
		return err
	}

    config_file := path.Join(config_folder, filename)

	content, err := ioutil.ReadFile(config_file)
	if err != nil {
		return errors.New("Could not find config file at " + config_file)
	}

	if _, err := toml.Decode(string(content), conf); err != nil {
		return err
	}

	return nil
}

func GetServerConfig() (Config, error) {
    var conf Config
    if err := get_config(&conf, "config.toml"); err != nil {
        return conf, err
    }
    return conf, nil
}
