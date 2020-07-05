package utils

import "regexp"

var (
	MentionRegex = regexp.MustCompile(`<@\d{18}>`)
	UserIDRegex  = regexp.MustCompile(`\d{18}`)
)
