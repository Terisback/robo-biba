package utils

import (
	"fmt"
	"strings"
)

// Arguments holds all of the fields of the command
// First element always command itself
type Arguments []string

type ArgumentsKey string

const (
	ArgumentsContextKey ArgumentsKey = "args"
)

func ParseArguments(content, selfID string, prefixes ...string) (args Arguments, ok bool) {
	// Normalize
	content = strings.ToLower(content)
	content = strings.TrimLeft(content, " ")
	selfMention := fmt.Sprintf(`<@%s>`, selfID)

	// Find the prefixes and trim them
	if prefix, ok := hasOneOfThePrefixes(content, prefixes...); ok {
		// Usual prefix
		content = strings.TrimPrefix(content, prefix)
	} else if strings.HasPrefix(content, selfMention) {
		// Mention prefix
		content = strings.TrimPrefix(content, selfMention)
	} else {
		return nil, false
	}

	// Returning command arguments
	args, ok = strings.Fields(content), true
	return
}

func hasOneOfThePrefixes(content string, prefixes ...string) (prefix string, ok bool) {
	for _, prefix = range prefixes {
		prefix = strings.ToLower(prefix)
		prefix = strings.TrimSpace(prefix)
		if ok = strings.HasPrefix(content, prefix); ok {
			return
		}
	}

	return "", false
}
