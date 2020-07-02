package command

import "strings"

// Aliases compares text and aliases, then return it is equal or not
func Aliases(text string, aliases ...string) (ok bool) {
	// Normalize
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)

	for _, alias := range aliases {
		// Normalize
		alias = strings.ToLower(alias)
		alias = strings.TrimSpace(alias)

		// Compare
		if text == alias {
			return true
		}
	}
	return false
}
