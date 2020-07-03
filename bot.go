package main

import (
	"context"
	"log"

	"github.com/Terisback/robo-biba/commands"
	"github.com/Terisback/robo-biba/middleware"
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
		commands.Online)

	dg.On(disgord.EvtMessageCreate,
		mdl.FilterBotMessages,
		mdl.FilterPrefix(prefix),
		mdl.FilterAliases("w", "when", "к", "когда"),
		commands.When)

	log.Println("Bot is started!")
}
