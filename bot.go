package main

import (
	"context"
	"log"

	"github.com/Terisback/robo-biba/commands"
	"github.com/Terisback/robo-biba/middleware"
	"github.com/andersfylling/disgord"
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
		mdl.FilterCommand(middleware.CommandOptions{
			Prefixes: []string{config.Prefix},
			Aliases:  []string{"o", "online", "о", "онлайн"},
		}),
		commands.Online)

	dg.On(disgord.EvtMessageCreate,
		mdl.FilterBotMessages,
		mdl.FilterCommand(middleware.CommandOptions{
			Prefixes: []string{config.Prefix},
			Aliases:  []string{"w", "when", "к", "когда"},
		}),
		commands.When)

	log.Println("Bot is started!")
}
