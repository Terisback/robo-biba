package commands

import (
	"context"
	"log"
	"strconv"

	"github.com/Terisback/robo-biba/utils"
	"github.com/andersfylling/disgord"
)

const roleOfActivePeople = 665980888869371955

func Online(session disgord.Session, event *disgord.MessageCreate) {
	args, err := utils.GetArgsFromContext(event.Ctx)
	if err != nil {
		log.Println(err)
		return
	}

	// Role ID for monitor online of the role into second column
	var roleID uint64

	if len(args) == 2 {
		roleID, err = strconv.ParseUint(args[1], 10, 64)
		if err != nil {
			log.Println(err)
			roleID = roleOfActivePeople
		}
	} else {
		roleID = roleOfActivePeople
	}

	g, err := session.GetGuild(context.Background(), event.Message.GuildID)
	if err != nil {
		log.Println(err)
		return
	}

	doSecondColumn := true

	role, err := g.Role(disgord.NewSnowflake(roleID))
	if err != nil {
		doSecondColumn = false
	}

	activeOnline := 0
	online := 0
	for _, presence := range g.Presences {
		if presence.Status != "offline" {
			online += 1

			u, err := g.Member(presence.User.ID)
			if err != nil {
				log.Println(err)
				return
			}

			if doSecondColumn {
				for _, r := range u.Roles {
					if r == role.ID {
						activeOnline += 1
					}
				}
			}
		}
	}

	allActive := 0
	if doSecondColumn {
		for _, member := range g.Members {
			for _, r := range member.Roles {
				if r == role.ID {
					allActive += 1
				}
			}
		}
	}

	var embeds []*disgord.EmbedField

	embeds = append(embeds, &disgord.EmbedField{
		Name:   "Общак",
		Value:  "Всего: " + strconv.Itoa(int(g.MemberCount)) + "\n" + "Онлайн: " + strconv.Itoa(online),
		Inline: true,
	},
	)

	if doSecondColumn {
		embeds = append(embeds, &disgord.EmbedField{
			Name:   role.Name,
			Value:  "Всего: " + strconv.Itoa(allActive) + "\n" + "Онлайн: " + strconv.Itoa(activeOnline),
			Inline: true,
		},
		)
	}

	_, err = event.Message.Reply(context.Background(), session,
		disgord.Embed{
			Title:  "Онлайн " + g.Name,
			Fields: embeds,
		},
	)
	if err != nil {
		log.Println(err)
		return
	}
}
