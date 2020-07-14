package commands

import (
	"context"
	"log"

	"github.com/Terisback/robo-biba/middleware"
	"github.com/Terisback/robo-biba/utils"
	"github.com/andersfylling/disgord"
)

func When(session disgord.Session, event *disgord.MessageCreate) {
	command, err := middleware.GetCommandFromContext(event.Ctx)
	if err != nil {
		log.Println(err)
		return
	}

	args := command.Arguments

	if len(args) == 1 {
		_, err := event.Message.Reply(context.Background(), session, whenEmbed(event.Message.Member))
		if err != nil {
			log.Println(err)
		}
		return
	}

	if len(args) < 2 {
		embed := disgord.Embed{}
		embed.Description = "`!когда <id или упоминание>` - Узнать когда пользователь зашёл на сервер"
		_, err := event.Message.Reply(context.Background(), session, &embed)
		if err != nil {
			log.Println(err)
		}
		return
	}

	if _, ok := args[1].GetID(); ok {
		var member *disgord.Member
		member, _ = utils.GetMemberFromArg(session, event.Message.GuildID, args[1].String())
		_, err = event.Message.Reply(context.Background(), session, whenEmbed(member))
		if err != nil {
			log.Println(err)
		}
		return
	}

	embed := disgord.Embed{}
	embed.Description = "`!когда <id или упоминание>` - Узнать когда пользователь зашёл на сервер"
	_, err = event.Message.Reply(context.Background(), session, &embed)
	if err != nil {
		log.Println(err)
	}
}

func whenEmbed(member *disgord.Member) *disgord.Embed {
	embed := disgord.Embed{}
	embed.Color = utils.GetIntColor(utils.DefaultEmbedColor)

	nickname := member.Nick

	if nickname == "" {
		nickname = member.User.Username
	}

	userJoinedAtDate := member.JoinedAt.Time
	embed.Description = "**" + nickname + userJoinedAtDate.Format("** зашёл на сервер 02.01.2006 в 15:04:05 по Москве")
	return &embed
}
