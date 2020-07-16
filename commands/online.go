package commands

import (
	"context"
	"log"
	"strconv"

	"github.com/Terisback/robo-biba/internal/utils"
	"github.com/Terisback/robo-biba/middleware"
	"github.com/andersfylling/disgord"
)

const roleOfActivePeople = 665980888869371955

func Online(session disgord.Session, event *disgord.MessageCreate) {
	command, err := middleware.GetCommandFromContext(event.Ctx)
	if err != nil {
		log.Println(err)
		return
	}

	var (
		args      = command.Arguments
		roleID    uint64
		ok        bool
		roleField bool = true
	)

	if len(args) == 2 {
		roleID, ok = args[1].GetID()
		if !ok {
			roleID = roleOfActivePeople
		}
	} else {
		roleID = roleOfActivePeople
	}

	guild, err := session.GetGuild(context.Background(), event.Message.GuildID)
	if err != nil {
		log.Println(err)
		return
	}

	role, err := guild.Role(disgord.NewSnowflake(roleID))
	if err != nil {
		roleField = false
	}

	everyoneOnline := 0
	roleOnline := 0
	for _, presence := range guild.Presences {
		if presence.Status != "offline" {
			everyoneOnline += 1

			if roleField {
				u, err := guild.Member(presence.User.ID)
				if err != nil {
					log.Println(err)
					return
				}

				for _, r := range u.Roles {
					if r == role.ID {
						roleOnline += 1
					}
				}
			}
		}
	}

	everyoneCount := guild.MemberCount
	roleCount := 0
	if roleField {
		for _, member := range guild.Members {
			for _, r := range member.Roles {
				if r == role.ID {
					roleCount += 1
				}
			}
		}
	}

	var embeds []*disgord.EmbedField

	embeds = append(embeds, &disgord.EmbedField{
		Name:   "Общак",
		Value:  "Всего: " + strconv.Itoa(int(everyoneCount)) + "\n" + "Онлайн: " + strconv.Itoa(everyoneOnline),
		Inline: true,
	},
	)

	if roleField {
		embeds = append(embeds, &disgord.EmbedField{
			Name:   role.Name,
			Value:  "Всего: " + strconv.Itoa(roleCount) + "\n" + "Онлайн: " + strconv.Itoa(roleOnline),
			Inline: true,
		})
	}

	_, err = event.Message.Reply(context.Background(), session,
		disgord.Embed{
			Title:  "Онлайн " + guild.Name,
			Fields: embeds,
			Color:  utils.GetIntColor(utils.DefaultEmbedColor),
		},
	)
	if err != nil {
		log.Println(err)
		return
	}
}
