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
	command, err := middleware.GetCommandFromContext(event.Ctx)
	if err != nil {
		log.Println(err)
		return
	}

	var (
		embed    = disgord.Embed{}
		nickname string
		args     = command.Arguments
	)

	if len(args) == 1 {
		nickname = event.Message.Member.Nick

		if nickname == "" {
			nickname = event.Message.Author.Username
		}

		userJoinedAtDate := event.Message.Member.JoinedAt
		embed.Description = "**" + nickname + userJoinedAtDate.Format("** зашел на сервер 02.01.2006 в 15:04:05 по Москве")
		_, err := event.Message.Reply(context.Background(), session, &embed)
		if err != nil {
			log.Println(err)
		}
		return
	}

	if len(args) < 2 {
		embed.Description = "`!когда <id или упоминание>` - Узнать когда пользователь зашёл на сервер"
		_, err := event.Message.Reply(context.Background(), session, &embed)
		if err != nil {
			log.Println(err)
		}
		return
	}

	if _, ok := args[1].GetID(); ok {
		member, _ = utils.GetMemberFromArg(session, event.Message.GuildID, args[1].String())

		nickname = member.Nick

		if nickname == "" {
			nickname = member.User.Username
		}

		userJoinedAtDate := member.JoinedAt.Time
		embed.Description = "**" + nickname + userJoinedAtDate.Format("** зашёл на сервер 02.01.2006 в 15:04:05 по Москве")
		_, err = event.Message.Reply(context.Background(), session, &embed)
		if err != nil {
			log.Println(err)
		}
		return
	}

	// Return created date from Author ID
	if utils.Aliases(args[1].String(), "я", "мой", "me", "my") {
		nickname = event.Message.Member.Nick

		if nickname == "" {
			nickname = event.Message.Author.Username
		}

		userJoinedAtDate := event.Message.Member.JoinedAt
		embed.Description = "**" + nickname + userJoinedAtDate.Format("** зашел на сервер 02.01.2006 в 15:04:05 по Москве")
		_, err := event.Message.Reply(context.Background(), session, &embed)
		if err != nil {
			log.Println(err)
		}
		return
	}

	embed.Description = "`!когда <id или упоминание>` - Узнать когда пользователь зашёл на сервер"
	_, err = event.Message.Reply(context.Background(), session, &embed)
	if err != nil {
		log.Println(err)
	}
}
