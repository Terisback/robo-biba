package commands

import "strconv"

const (
	defaultEmbedColor = "3498db"
)

func getIntColor(embColor string) int {
	color, err := strconv.ParseUint(embColor, 16, 64)
	if err != nil {
		return 0
	}
	return int(color)
}
