package commands

import (
	"context"
	"fmt"
	"log"

	"github.com/Terisback/robo-biba/internal/boolgen"
	"github.com/Terisback/robo-biba/internal/storage"
	"github.com/Terisback/robo-biba/internal/utils"
	"github.com/Terisback/robo-biba/middleware"
	"github.com/andersfylling/disgord"
)

const (
	cfWinColor    = "66ed4e"
	cfLoseColor   = "f55847"
	cfNeutralCoin = "https://cdn.discordapp.com/emojis/518875768814829568.png?v=1"
	cfWinCoin     = "https://cdn.discordapp.com/emojis/518875768814829568.png?v=1"
	cfLoseCoin    = "https://cdn.discordapp.com/emojis/518875812913610754.png?v=1"
	cfMessage     = "**%s**\n__Ставка:__ **%d** <:rgd_coin_rgd:518875768814829568>\n__Баланс:__ **%d** <:rgd_coin_rgd:518875768814829568>"
	cfHelp        = "Для игры напишите `!флип <сумма>`"
)

var (
	bg = boolgen.New()
)

func Coinflip(session disgord.Session, event *disgord.MessageCreate) {
	command, err := middleware.GetCommandFromContext(event.Ctx)
	if err != nil {
		log.Println(err)
		return
	}

	args := command.Arguments

	if len(args) < 2 {
		cfSendHelpMessage(session, event)
		return
	}

	if bet, ok := args[1].GetNumeric(); ok {
		guildID := event.Message.GuildID.String()
		userID := event.Message.Author.ID.String()

		balance, err := storage.Balance(guildID, userID)
		if err != nil {
			return
		}

		if bet < 0 {
			bet = -bet
		}

		if balance-bet < 0 {
			cfSendNotEnoughMessage(session, event)
			return
		}

		win := bg.RandBool()

		msg := cfSendGameStartMessage(session, event, bet, balance)

		if win {
			cfSendWinMessage(session, event, msg, bet)
		} else {
			cfSendLoseMessage(session, event, msg, bet)
		}

	}

	cfSendHelpMessage(session, event)
}

func cfSendGameStartMessage(session disgord.Session, event *disgord.MessageCreate, bet, balance int) (msg *disgord.Message) {
	nickname := event.Message.Member.Nick

	if nickname == "" {
		nickname = event.Message.Author.Username
	}

	avatarURL, err := event.Message.Author.AvatarURL(64, false)
	if err != nil {
		return
	}

	embed := utils.GetDefaultEmbed()
	embed.Author = &disgord.EmbedAuthor{IconURL: avatarURL, Name: fmt.Sprintf("%s подбросил монетку", nickname)}
	embed.Description = fmt.Sprintf(cfMessage, "ПОДБРАСЫВАЕМ...", bet, balance)
	embed.Thumbnail = &disgord.EmbedThumbnail{URL: cfNeutralCoin}

	msg, err = event.Message.Reply(context.Background(), session, embed)
	if err != nil {
		return &disgord.Message{}
	}

	return msg
}

func cfSendWinMessage(session disgord.Session, event *disgord.MessageCreate, msg *disgord.Message, bet int) {
	guildID := event.Message.GuildID.String()
	userID := event.Message.Author.ID.String()

	balance, err := storage.AddCurrency(guildID, userID, bet)
	if err != nil {
		log.Println("Economy go brrrrrr")
		return
	}

	embed := &disgord.Embed{}
	embed.Description = fmt.Sprintf(cfMessage, "ПОБЕДА!", bet, balance)
	embed.Thumbnail = &disgord.EmbedThumbnail{URL: cfWinCoin}
	embed.Color = utils.GetIntColor(cfWinColor)

	_, _ = session.SetMsgEmbed(context.Background(), msg.ChannelID, msg.ID, embed)
}

func cfSendLoseMessage(session disgord.Session, event *disgord.MessageCreate, msg *disgord.Message, bet int) {
	guildID := event.Message.GuildID.String()
	userID := event.Message.Author.ID.String()

	balance, err := storage.SubCurrency(guildID, userID, bet)
	if err != nil {
		log.Println("Economy go brrrrrr")
		return
	}

	embed := &disgord.Embed{}
	embed.Description = fmt.Sprintf(cfMessage, "ПОРАЖЕНИЕ", bet, balance)
	embed.Thumbnail = &disgord.EmbedThumbnail{URL: cfLoseCoin}
	embed.Color = utils.GetIntColor(cfLoseColor)

	_, _ = session.SetMsgEmbed(context.Background(), msg.ChannelID, msg.ID, embed)
}

func cfSendHelpMessage(session disgord.Session, event *disgord.MessageCreate) {
	embed := utils.GetDefaultEmbed()
	embed.Description = cfHelp
	_, _ = event.Message.Reply(context.Background(), session, embed)
}

func cfSendNotEnoughMessage(session disgord.Session, event *disgord.MessageCreate) {
	nickname := event.Message.Member.Nick

	if nickname == "" {
		nickname = event.Message.Author.Username
	}

	avatarURL, err := event.Message.Author.AvatarURL(64, false)
	if err != nil {
		return
	}

	embed := utils.GetDefaultEmbed()
	embed.Author = &disgord.EmbedAuthor{IconURL: avatarURL, Name: fmt.Sprintf("%s вы не можете играть на сумму превышающую ваш баланс", nickname)}
	_, _ = event.Message.Reply(context.Background(), session, embed)
}
