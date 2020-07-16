package utils

import (
	"github.com/andersfylling/disgord"
)

func GetDefaultEmbed() *disgord.Embed {
	embed := disgord.Embed{}
	embed.Color = GetIntColor(DefaultEmbedColor)
	return &embed
}
