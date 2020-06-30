package handlers

import (
	"context"
	"log"
	"strconv"

	"github.com/andersfylling/disgord"
)

const roleOfActivePeople = 665980888869371955

func Online(s disgord.Session, e *disgord.MessageCreate) {
	g, err := s.GetGuild(context.Background(), e.Message.GuildID)
	if err != nil {
		log.Println(err)
		return
	}

	activeOnline := 0
	online := 0
	for _, i := range g.Presences {
		if i.Status != "offline" {
			online += 1

			u, err := g.Member(i.User.ID)
			if err != nil {
				log.Println(err)
				return
			}

			for _, r := range u.Roles {
				if r == disgord.NewSnowflake(roleOfActivePeople) {
					activeOnline += 1
				}
			}
		}
	}

	allActive := 0
	for _, i := range g.Members {
		for _, r := range i.Roles {
			if r == disgord.NewSnowflake(roleOfActivePeople) {
				allActive += 1
			}
		}
	}

	_, err = e.Message.Reply(context.Background(), s,
		disgord.Embed{
			Title: "Онлайн " + g.Name,
			Fields: []*disgord.EmbedField{
				&disgord.EmbedField{
					Name:   "Общак",
					Value:  "Всего: " + strconv.Itoa(int(g.MemberCount)) + "\n" + "Онлайн: " + strconv.Itoa(online),
					Inline: true,
				},
				&disgord.EmbedField{
					Name:   "Активных",
					Value:  "Всего: " + strconv.Itoa(allActive) + "\n" + "Онлайн: " + strconv.Itoa(activeOnline),
					Inline: true,
				},
			},
		})
	if err != nil {
		log.Println(err)
		return
	}
}
