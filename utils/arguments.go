package utils

import (
	"context"
	"errors"
	"strings"
)

// Arguments holds all of the fields of the command
// First element always command itself
type Arguments []string

// Parse determines if the normalized command has a normalized prefix; in this case returns the normalized arguments of the command
func Parse(selfID, prefix, content string) (arg Arguments, ok bool) {
	// Prepare for next step
	content = strings.ToLower(content)
	content = strings.TrimLeft(content, " ")
	prefix = strings.ToLower(prefix)

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

// Allows you to get arguments from context after FilterPrefix middleware
func GetArgsFromContext(ctx context.Context) (Arguments, error) {
	args := ctx.Value("args").(Arguments)

	if args == nil || len(args) == 0 {
		return nil, errors.New("Args is nil or zero length")
	}

	return args, nil
}
