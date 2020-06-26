package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Terisback/robo-biba/internal/handlers"
	"github.com/bwmarrin/discordgo"
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
		log.Fatal(err)
	}
}

func main() {
	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = session.Open()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Bot started...")
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-sc
	}()

	// This is not okay, TODO: Make internal router
	session.AddHandler(handlers.Online)
}
