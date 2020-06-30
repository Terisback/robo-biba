package handlers

import (
	"context"
	"errors"
	"log"

	"github.com/Terisback/robo-biba/internal/command"
	"github.com/andersfylling/disgord"
)

func When(s disgord.Session, e *disgord.MessageCreate) {
	// Getting mentions without mention of the bot
	var (
		user *disgord.User
		err  error
	)
	if user, err = s.GetCurrentUser(context.Background()); err != nil {
		log.Println(errors.New("Unable to fetch info about the bot instance"))
		return
	}
	mentions := e.Message.Mentions
	n := 0
	for _, m := range mentions {
		if m.ID != user.ID {
			mentions[n] = m
			n++
		}
	}
	mentions = mentions[:n]

	args, err := command.GetArgsFromContext(e.Ctx)
	if err != nil {
		log.Println(err)
		return
	}

	// Return time when the server was created
	if len(mentions) == 0 && len(args) == 1 {
		serverCreatedDate := e.Message.GuildID.Date()
		_, err := e.Message.Reply(context.Background(), s,
			disgord.Message{
				Content: serverCreatedDate.Format("Сервер был создан 2.01.2006 в 15:04:05 по Москве"),
			},
		)
		if err != nil {
			log.Println(err)
		}
		return
	}

	// Return created date from Author ID
	if command.Aliases(args[1], "я", "мой", "me", "my") {
		userCreatedDate := e.Message.Author.ID.Date()
		_, err := e.Message.Reply(context.Background(), s,
			disgord.Message{
				Content: "**" + e.Message.Author.Username + userCreatedDate.Format("** был создан 2.01.2006 в 15:04:05 по Москве"),
			},
		)
		if err != nil {
			log.Println(err)
		}
		return
	}

	if len(mentions) == 0 {
		// TODO: Reply with help
		return
	}

	// Return created date from first mention ID
	if command.Aliases(args[1], "создан", "зарегался", "registered", "reg", "created") {
		userCreatedDate := mentions[0].ID.Date()
		_, err := e.Message.Reply(context.Background(), s,
			disgord.Message{
				Content: "**" + mentions[0].Username + userCreatedDate.Format("** был создан 2.01.2006 в 15:04:05 по Москве"),
			},
		)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if command.Aliases(args[1], "зашёл", "зашел", "присоединился", "joined", "join") {
		member, err := s.GetMember(context.Background(), e.Message.GuildID, mentions[0].ID)
		if err != nil {
			log.Println(err)
			return
		}
		userCreatedDate := member.JoinedAt.Time
		_, err = e.Message.Reply(context.Background(), s,
			disgord.Message{
				Content: "**" + member.Nick + userCreatedDate.Format("** зашёл на сервер 2.01.2006 в 15:04:05 по Москве"),
			},
		)
		if err != nil {
			log.Println(err)
			return
		}
	}

	{
		member, err := s.GetMember(context.Background(), e.Message.GuildID, mentions[0].ID)
		if err != nil {
			log.Println(err)
			return
		}
		userCreatedDate := member.JoinedAt.Time
		_, err = e.Message.Reply(context.Background(), s,
			disgord.Message{
				Content: "**" + member.Nick + userCreatedDate.Format("** зашёл на сервер 2.01.2006 в 15:04:05 по Москве"),
			},
		)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
