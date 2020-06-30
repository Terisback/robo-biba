package main

import (
	"log"

	"github.com/codingconcepts/env"
)

type Config struct {
	Token string `env:"TOKEN" required:"true"`
}

var (
	config Config
)

func init() {
	if err := env.Set(&config); err != nil {
		log.Fatalln("Error caused by trying to get environment variables,", err)
	}
}
