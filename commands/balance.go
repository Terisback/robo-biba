package commands

import (
	"context"
	"fmt"

	"github.com/Terisback/robo-biba/internal/storage"
	"github.com/andersfylling/disgord"
)

const (
	blDesc = "Ваш баланс: **%d** <:rgd_coin_rgd:518875768814829568>"
)

func Balance(session disgord.Session, event *disgord.MessageCreate) {
	var nickname string = event.Message.Member.Nick

	if nickname == "" {
		nickname = event.Message.Author.Username
	}

	avatarURL, err := event.Message.Author.AvatarURL(64, false)
	if err != nil {
		return
	}

	balance, err := storage.Balance(event.Message.GuildID.String(), event.Message.Author.ID.String())
	if err != nil {
		return
	}

	embed := disgord.Embed{}
	embed.Author = &disgord.EmbedAuthor{IconURL: avatarURL, Name: nickname}
	embed.Description = fmt.Sprintf(blDesc, balance)

	session.SendMsg(context.Background(), event.Message.ChannelID, &embed)
}