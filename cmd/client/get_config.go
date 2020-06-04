package main

import (
	"errors"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

func get_client_config(conf *Config) error {
	pwd, pok := os.LookupEnv("PWD")
	if !pok {
		return errors.New("Could not get $PWD")
	}

	file := path.Join(pwd, "./cmd/client/config_client.toml")

	if _, err := toml.DecodeFile(file, conf); err != nil {
		return err
	}

	return nil
}
