package utils

import (
	"errors"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// const ActiveServerID = "701878021421793330"

// Checks if the message is valid, not empty, starts with '!', wasn't sent by the bot
func CheckValid(s *discordgo.Session, e *discordgo.MessageCreate) (err error) {
	message := e.Message
	content := message.Content

	if message.Author.Bot {
		return errors.New("Bot message")
	}

	content = strings.TrimSpace(content)

	if content == "" {
		return errors.New("Content is empty")
	}

	if !strings.HasPrefix(content, "!") {
		return errors.New("No prefix")
	}

	// if e.GuildID != ActiveServerID {
	// 	return errors.New("Isn't active server (temporary, will be deleted after implementing router)")
	// }

	return nil
}

// Checks similarity of the message content to aliases
func CheckCommand(content string, aliases ...string) bool {
	content = strings.ToLower(strings.TrimSpace(content))
	for _, alias := range aliases {
		alias = strings.ToLower(strings.TrimSpace(alias))
		if strings.HasPrefix(content, "!"+alias) {
			return true
		}
	}
	return false
}
