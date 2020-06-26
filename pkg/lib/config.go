package lib

import (
	"errors"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"path"
)

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

func GetConfig(conf interface{}, filename string, config_file string) error {
    if config_file == "" {
        config_folder, err := get_config_folder()
        if err != nil {
            return err
        }

        config_file = path.Join(config_folder, filename)
    }

	content, err := ioutil.ReadFile(config_file)
	if err != nil {
		return errors.New("Could not find config file at " + config_file)
	}

	if _, err := toml.Decode(string(content), conf); err != nil {
		return err
	}

	return nil
}
