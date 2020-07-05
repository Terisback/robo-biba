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
	var mentionedID uint64
	var thereIsMention bool
	args, err := middleware.GetArgsFromContext(event.Ctx)
	if err != nil {
		log.Println(err)
		return
	}

	if len(args) == 1 {
		serverCreatedDate := event.Message.GuildID.Date()
		_, err := event.Message.Reply(context.Background(), session,
			disgord.Message{
				Content: serverCreatedDate.Format("Сервер был создан 02.01.2006 в 15:04:05 по Москве"),
			},
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

	if mentionedID, thereIsMention = utils.GetIDFromArg(args[1]); thereIsMention {
		member, err = session.GetMember(context.Background(), event.Message.GuildID, disgord.NewSnowflake(mentionedID))
		if err != nil {
			log.Println(err)
			return
		}
	}

	if thereIsMention {
		userCreatedDate := member.JoinedAt.Time
		_, err = event.Message.Reply(context.Background(), session,
			disgord.Message{
				Content: "**" + member.User.Username + userCreatedDate.Format("** зашёл на сервер 02.01.2006 в 15:04:05 по Москве"),
			},
		)
		if err != nil {
			log.Println(err)
			return
		}
	}

	// Return created date from Author ID
	if utils.Aliases(args[1], "я", "мой", "me", "my") {
		userCreatedDate := event.Message.Author.ID.Date()
		_, err := event.Message.Reply(context.Background(), session,
			disgord.Message{
				Content: "**" + event.Message.Author.Username + userCreatedDate.Format("** был создан 02.01.2006 в 15:04:05 по Москве"),
			},
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

	if mentionedID, thereIsMention = utils.GetIDFromArg(args[2]); thereIsMention {
		member, err = session.GetMember(context.Background(), event.Message.GuildID, disgord.NewSnowflake(mentionedID))
		if err != nil {
			log.Println(err)
			return
		}
	}

	// Return created date from first mention ID
	if thereIsMention && utils.Aliases(args[1], "создан", "зарегался", "registered", "reg", "created") {
		userCreatedDate := member.User.ID.Date()
		_, err := event.Message.Reply(context.Background(), session,
			disgord.Message{
				Content: "**" + member.User.Username + userCreatedDate.Format("** был создан 02.01.2006 в 15:04:05 по Москве"),
			},
		)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if thereIsMention && utils.Aliases(args[1], "зашёл", "зашел", "присоединился", "joined", "join") {
		userCreatedDate := member.JoinedAt.Time
		_, err = event.Message.Reply(context.Background(), session,
			disgord.Message{
				Content: "**" + member.User.Username + userCreatedDate.Format("** зашёл на сервер 02.01.2006 в 15:04:05 по Москве"),
			},
		)
		if err != nil {
			log.Println(err)
			return
		}
	}

	// Return help
}
