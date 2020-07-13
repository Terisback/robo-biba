package main

import (
	"log"

	"github.com/codingconcepts/env"
)

type Config struct {
	Token   string `env:"TOKEN" required:"true"`
	Prefix  string `env:"BOT_PREFIX" default:"!"`
	MongoDB string `env:"MONGODB" required:"true"`
}

var (
	config Config
)

func init() {
	if err := env.Set(&config); err != nil {
		log.Fatalln("Error caused by trying to get environment variables,", err)
	}
}
