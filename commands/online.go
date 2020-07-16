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

	var embeds []*disgord.EmbedField

	embeds = append(embeds, onlineEveryoneField(guild))

	if roleField {
		embeds = append(embeds, onlineRoleField(guild, role))
	}

	embed := utils.GetDefaultEmbed()
	embed.Title = "Онлайн " + guild.Name
	embed.Fields = embeds

	_, err = event.Message.Reply(context.Background(), session, embed)
	if err != nil {
		log.Println(err)
		return
	}
}

func onlineEveryoneField(guild *disgord.Guild) *disgord.EmbedField {
	var everyoneOnline int
	for _, presence := range guild.Presences {
		if presence.Status != "offline" {
			everyoneOnline += 1
		}
	}

	everyoneCount := guild.MemberCount

	return &disgord.EmbedField{
		Name:   "Общак",
		Value:  "Всего: " + strconv.Itoa(int(everyoneCount)) + "\n" + "Онлайн: " + strconv.Itoa(everyoneOnline),
		Inline: true,
	}
}

func onlineRoleField(guild *disgord.Guild, role *disgord.Role) *disgord.EmbedField {
	roleOnline := 0
	for _, presence := range guild.Presences {
		if presence.Status != "offline" {
			u, err := guild.Member(presence.User.ID)
			if err != nil {
				continue
			}

			for _, r := range u.Roles {
				if r == role.ID {
					roleOnline += 1
				}
			}

		}
	}

	roleCount := 0
	for _, member := range guild.Members {
		for _, r := range member.Roles {
			if r == role.ID {
				roleCount += 1
			}
		}
	}

	return &disgord.EmbedField{
		Name:   role.Name,
		Value:  "Всего: " + strconv.Itoa(roleCount) + "\n" + "Онлайн: " + strconv.Itoa(roleOnline),
		Inline: true,
	}
}
