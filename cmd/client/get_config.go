package main

import (
	"github.com/kyazdani42/taskmaster/pkg/lib"
	"log"
)

func getClientConfig(c *Config) {
	filename := "config_client.toml"
	err := lib.GetConfig(c, filename, "")
	if err != nil {
		log.Fatalln(err)
	}
}
