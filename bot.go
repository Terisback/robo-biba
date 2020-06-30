package main

import (
	"context"
	"log"

	"github.com/Terisback/robo-biba/internal/handlers"
	"github.com/Terisback/robo-biba/internal/middleware"
	"github.com/andersfylling/disgord"
)

const (
	prefix = "!"
)

func main() {
	dg := disgord.New(disgord.Config{
		BotToken: config.Token,
	})
	defer dg.StayConnectedUntilInterrupted(context.Background())

	mdl, err := middleware.New(dg)
	if err != nil {
		log.Fatalln(err)
	}

	dg.On(disgord.EvtMessageCreate,
		mdl.FilterBotMessages,
		mdl.FilterPrefix(prefix),
		mdl.FilterAliases("o", "online", "о", "онлайн"),
		handlers.Online)

	log.Println("Bot is started!")
}
