package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

func MustInit(path string) *Config {

	if path == "" {
		log.Fatal("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
