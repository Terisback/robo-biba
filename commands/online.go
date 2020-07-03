package commands

import (
	"context"
	"log"
	"strconv"

	"github.com/Terisback/robo-biba/utils"
	"github.com/andersfylling/disgord"
)

const roleOfActivePeople = 665980888869371955

func Online(s disgord.Session, e *disgord.MessageCreate) {
	var roleID uint64

	args, err := utils.GetArgsFromContext(e.Ctx)
	if err != nil {
		log.Println(err)
		return
	}

	if len(args) == 2 {
		roleID, err = strconv.ParseUint(args[1], 10, 64)
		if err != nil {
			log.Println(err)
			roleID = roleOfActivePeople
		}
	} else {
		roleID = roleOfActivePeople
	}

	g, err := s.GetGuild(context.Background(), e.Message.GuildID)
	if err != nil {
		log.Println(err)
		return
	}

	role, err := g.Role(disgord.NewSnowflake(roleID))
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
				if r == role.ID {
					activeOnline += 1
				}
			}
		}
	}

	allActive := 0
	for _, i := range g.Members {
		for _, r := range i.Roles {
			if r == role.ID {
				allActive += 1
			}
		}
	}

	_, err = e.Message.Reply(context.Background(), s,
		disgord.Embed{
			Title: "Онлайн " + g.Name,
			Fields: []*disgord.EmbedField{
				{
					Name:   "Общак",
					Value:  "Всего: " + strconv.Itoa(int(g.MemberCount)) + "\n" + "Онлайн: " + strconv.Itoa(online),
					Inline: true,
				},
				{
					Name:   role.Name,
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
