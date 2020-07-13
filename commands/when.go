package commands

import (
	"context"
	"log"

	"github.com/Terisback/robo-biba/middleware"
	"github.com/Terisback/robo-biba/utils"
	"github.com/andersfylling/disgord"
)

func When(session disgord.Session, event *disgord.MessageCreate) {
	var member *disgord.Member
	var thereIsMention bool
	command, err := middleware.GetCommandFromContext(event.Ctx)
	if err != nil {
		log.Println(err)
		return
	}

	args := command.Arguments

	if len(args) == 1 {
		serverCreatedDate := event.Message.GuildID.Date()
		_, err := event.Message.Reply(context.Background(), session,
			serverCreatedDate.Format("Сервер был создан 02.01.2006 в 15:04:05 по Москве"),
		)
		if err != nil {
			log.Println(err)
		}
		return
	}

	if len(args) < 2 {
		// TODO: Return help
		return
	}

	var nickname string
	thereIsMention = false

	member, err = session.GetMember(context.Background(), event.Message.GuildID, event.Message.Author.ID)
	if err != nil {
		return
	}
	nickname = member.Nick

	if nickname == "" {
		nickname = member.User.Username
	}

	if _, ok := args[1].GetID(); ok {
		member, thereIsMention = utils.GetMemberFromArg(session, event.Message.GuildID, args[1].String())

		nickname = member.Nick

		if nickname == "" {
			nickname = member.User.Username
		}
	}

	if thereIsMention {
		userCreatedDate := member.JoinedAt.Time
		_, err = event.Message.Reply(context.Background(), session,
			"**"+nickname+userCreatedDate.Format("** зашёл на сервер 02.01.2006 в 15:04:05 по Москве"),
		)
		if err != nil {
			log.Println(err)
			return
		}
	}

	// Return created date from Author ID
	if utils.Aliases(args[1].String(), "я", "мой", "me", "my") {
		userCreatedDate := event.Message.Author.ID.Date()
		_, err := event.Message.Reply(context.Background(), session,
			"**"+nickname+userCreatedDate.Format("** был создан 02.01.2006 в 15:04:05 по Москве"),
		)
		if err != nil {
			log.Println(err)
		}
		return
	}

	if len(args) < 3 {
		// TODO: Return help
		return
	}

	thereIsMention = false

	if _, ok := args[2].GetID(); ok {
		member, thereIsMention = utils.GetMemberFromArg(session, event.Message.GuildID, args[2].String())

		nickname = member.Nick

		if nickname == "" {
			nickname = member.User.Username
		}
	}

	if thereIsMention && utils.Aliases(args[1].String(), "создан", "зарегался", "registered", "reg", "created") {
		userCreatedDate := member.User.ID.Date()
		_, err := event.Message.Reply(context.Background(), session,
			"**"+nickname+userCreatedDate.Format("** был создан 02.01.2006 в 15:04:05 по Москве"),
		)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if thereIsMention && utils.Aliases(args[1].String(), "зашёл", "зашел", "присоединился", "joined", "join") {
		userCreatedDate := member.JoinedAt.Time
		_, err = event.Message.Reply(context.Background(), session,
			"**"+nickname+userCreatedDate.Format("** зашёл на сервер 02.01.2006 в 15:04:05 по Москве"),
		)
		if err != nil {
			log.Println(err)
			return
		}
	}

	// Return help
}
