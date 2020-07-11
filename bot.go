package main

import (
	"context"
	"log"

	"github.com/Terisback/robo-biba/commands"
	"github.com/Terisback/robo-biba/economy"
	"github.com/Terisback/robo-biba/middleware"
	"github.com/andersfylling/disgord"
)

func main() {
	dg := disgord.New(disgord.Config{
		BotToken: config.Token,
	})
	defer func() {
		dg.StayConnectedUntilInterrupted(context.Background())
		commands.GiftCacheSave()
		economy.Close()
	}()

	mdl, err := middleware.New(dg)
	if err != nil {
		log.Fatalln(err)
	}

	dg.On(disgord.EvtMessageCreate,
		mdl.FilterBotMessages,
		mdl.FilterCommand(middleware.CommandOptions{
			Prefixes: []string{config.Prefix},
			Aliases:  []string{"h", "help", "х", "помощь", "хелп"},
		}),
		commands.Help)

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

	dg.On(disgord.EvtMessageCreate,
		mdl.FilterBotMessages,
		mdl.FilterCommand(middleware.CommandOptions{
			Prefixes: []string{config.Prefix},
			Aliases:  []string{"b", "balance", "б", "баланс"},
		}),
		commands.Balance)

	dg.On(disgord.EvtMessageCreate,
		mdl.FilterBotMessages,
		mdl.FilterCommand(middleware.CommandOptions{
			Prefixes: []string{config.Prefix},
			Aliases:  []string{"g", "gift", "г", "п", "подарок"},
		}),
		commands.Gift)

	dg.On(disgord.EvtMessageCreate,
		mdl.FilterBotMessages,
		mdl.FilterCommand(middleware.CommandOptions{
			Prefixes: []string{config.Prefix},
			Aliases:  []string{"f", "flip", "ф", "флип"},
		}),
		commands.Coinflip)

	log.Println("Bot is started!")
}
