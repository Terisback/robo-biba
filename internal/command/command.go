package command

import (
	"strings"
)

// Arguments holds all of the fields of the command
// First element always command itself
type Arguments []string

// Parse determines if the command has a prefix; in this case returns the arguments of the command
func Parse(selfID, prefix, content string) (arg Arguments, ok bool) {
	// Prepare for next step
	content = strings.TrimLeft(content, " ")

	// Find the prefixes and trim them
	// TODO: Make it with regex
	if strings.HasPrefix(content, prefix) {
		// Usual prefix
		content = strings.TrimPrefix(content, prefix)
	} else if strings.HasPrefix(content, "<@"+selfID+">") {
		// Mention prefix
		content = strings.TrimPrefix(content, "<@"+selfID+">")
	} else {
		ok = false
		return
	}

	// Returning command arguments
	arg = strings.Fields(content)
	ok = true
	return
}
