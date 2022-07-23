package configs

import (
	"github.com/caarlos0/env/v6"
	"log"
)

type config struct {
	ServerPort string `env:"SERVER_PORT"`
}

func Get() *config {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	return &cfg
}
