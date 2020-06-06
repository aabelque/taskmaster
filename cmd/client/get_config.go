package main

import (
	"log"

	"github.com/kyazdani42/taskmaster/pkg/lib"
)

func getClientConfig(c *Config) {
	filename := "config_client.toml"
	err := config.GetConfig(c, filename, "")
	if err != nil {
		log.Fatalln(err)
	}
}
