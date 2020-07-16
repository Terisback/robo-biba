package utils

import "strconv"

const (
	DefaultEmbedColor = "3498db"
)

func GetIntColor(embColor string) int {
	color, err := strconv.ParseUint(embColor, 16, 64)
	if err != nil {
		return 0
	}
	return int(color)
}
