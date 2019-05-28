package config;

import (
	"log"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Server string
	Database string
}

func (c *Config) Read() {
	// Leitura de config.toml e passando dados para a struct
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}