package commands

import (
	"context"
	"fmt"
	"log"

	"github.com/Terisback/robo-biba/internal/storage"
	"github.com/Terisback/robo-biba/middleware"
	"github.com/Terisback/robo-biba/utils"
	"github.com/andersfylling/disgord"
)

const (
	cfNeutralColor = "f5f05f"
	cfWinColor     = "66ed4e"
	cfLoseColor    = "f55847"
	cfNeutralCoin  = "https://cdn.discordapp.com/emojis/518875768814829568.png?v=1"
	cfWinCoin      = "https://cdn.discordapp.com/emojis/518875768814829568.png?v=1"
	cfLoseCoin     = "https://cdn.discordapp.com/emojis/518875812913610754.png?v=1"
	cfMessage      = "**%s**\n__Ставка:__ **%d** <:rgd_coin_rgd:518875768814829568>\n__Баланс:__ **%d** <:rgd_coin_rgd:518875768814829568>"
	cfHelp         = "Для игры напишите `!флип <сумма>`"
)

func Coinflip(session disgord.Session, event *disgord.MessageCreate) {
	command, err := middleware.GetCommandFromContext(event.Ctx)
	if err != nil {
		log.Println(err)
		return
	}

	var (
		args = command.Arguments
	)

	if len(args) < 2 {
		embed := disgord.Embed{}
		embed.Color = getIntColor(defaultEmbedColor)
		embed.Description = cfHelp
		_, err := session.SendMsg(context.Background(), event.Message.ChannelID, &embed)
		if err != nil {
			return
		}
		return
	}

	if bet, ok := args[1].GetNumeric(); ok {
		guildID := event.Message.GuildID.String()
		userID := event.Message.Author.ID.String()

		balance, err := storage.Balance(guildID, userID)
		if err != nil {
			_, err := event.Message.Reply(context.Background(), session, "Something went wrong...")
			if err != nil {
				return
			}
			return
		}

		if bet < 0 {
			bet = -bet
		}

		var nickname string = event.Message.Member.Nick

		if nickname == "" {
			nickname = event.Message.Author.Username
		}

		avatarURL, err := event.Message.Author.AvatarURL(64, false)
		if err != nil {
			return
		}

		if balance-bet < 0 {
			embed := disgord.Embed{}
			embed.Color = getIntColor(defaultEmbedColor)
			embed.Author = &disgord.EmbedAuthor{IconURL: avatarURL, Name: fmt.Sprintf("%s вы не можете играть на сумму превышающую ваш баланс", nickname)}
			_, err := event.Message.Reply(context.Background(), session, &embed)
			if err != nil {
				return
			}
			return
		}

		win := utils.RandBool()

		embed := disgord.Embed{}
		embed.Author = &disgord.EmbedAuthor{IconURL: avatarURL, Name: fmt.Sprintf("%s подбросил монетку", nickname)}
		embed.Description = fmt.Sprintf(cfMessage, "ПОДБРАСЫВАЕМ...", bet, balance)
		embed.Thumbnail = &disgord.EmbedThumbnail{URL: cfNeutralCoin}
		embed.Color = getIntColor(defaultEmbedColor)

		msg, err := event.Message.Reply(context.Background(), session, &embed)
		if err != nil {
			return
		}

		var newColor string
		var newCoin string
		var newMessage string

		if win {
			balance, err = storage.AddCurrency(guildID, userID, bet)
			if err != nil {
				log.Println("Economy go brrrrrr")
				return
			}
			newColor = cfWinColor
			newCoin = cfWinCoin
			newMessage = "ПОБЕДА!"
		} else {
			balance, err = storage.SubCurrency(guildID, userID, bet)
			if err != nil {
				log.Println("Economy go brrrrrr")
				return
			}
			newColor = cfLoseColor
			newCoin = cfLoseCoin
			newMessage = "ПОРАЖЕНИЕ"
		}

		embed.Description = fmt.Sprintf(cfMessage, newMessage, bet, balance)
		embed.Thumbnail = &disgord.EmbedThumbnail{URL: newCoin}
		embed.Color = getIntColor(newColor)

		_, err = session.SetMsgEmbed(context.Background(), msg.ChannelID, msg.ID, &embed)
		if err != nil {
			return
		}
	} else {
		embed := disgord.Embed{}
		embed.Description = cfHelp
		_, err := event.Message.Reply(context.Background(), session, &embed)
		if err != nil {
			return
		}
		return
	}
}
