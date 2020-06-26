package handlers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Terisback/robo-biba/internal/utils"
	"github.com/bwmarrin/discordgo"
)

const (
	onlineTitleF = "Онлайн %s"
)

func Online(s *discordgo.Session, e *discordgo.MessageCreate) {
	if err := utils.CheckValid(s, e); err != nil {
		return
	}

	if !utils.CheckCommand(e.Content, "online", "онлайн", "o", "о", "stat", "стат") {
		return
	}

	guild, err := s.Guild(e.GuildID)
	if err != nil {
		log.Println("Can't get guild with", e.GuildID, "id,", err)
		return
	}

	if err := s.RequestGuildMembers(e.GuildID, "", 0, true); err != nil {
		log.Println("Can't get presences from", e.GuildID, "id,", err)
		return
	}

	online := 0
	for _, member := range guild.Presences {
		if member.Status != discordgo.StatusOffline {
			online++
		}
	}

	if _, err := s.ChannelMessageSendEmbed(e.ChannelID, &discordgo.MessageEmbed{
		Title: fmt.Sprintf(onlineTitleF, guild.Name),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Всего: " + strconv.Itoa(guild.MemberCount),
				Value: "Онлайн: " + strconv.Itoa(online),
			},
		},
	}); err != nil {
		log.Println("Message wasn't sent, ", err)
		return
	}
}
