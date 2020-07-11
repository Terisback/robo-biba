package utils

import (
	"context"
	"regexp"
	"strconv"

	"github.com/andersfylling/disgord"
)

var (
	userIDRegex = regexp.MustCompile(`\d{18}`)
)

// Returns mentions slice without user with id
func GetMentionWithoutID(mentions []*disgord.User, id disgord.Snowflake) []*disgord.User {
	tmpMentions := make([]*disgord.User, len(mentions))

	for _, m := range mentions {
		if m.ID != id {
			tmpMentions = append(tmpMentions, m)
		}
	}

	return tmpMentions
}

// Returns id uint64 from argument string
func GetIDFromArg(arg string) (id uint64, ok bool) {
	var err error

	if userIDRegex.MatchString(arg) {
		userID := userIDRegex.FindString(arg)
		id, err = strconv.ParseUint(userID, 10, 64)
		if err != nil {
			return 0, false
		}
		return id, true
	}

	return 0, false
}

func GetMemberFromArg(s disgord.Session, guildID disgord.Snowflake, arg string) (member *disgord.Member, ok bool) {
	var id uint64
	id, ok = GetIDFromArg(arg)
	if !ok {
		return
	}

	var err error
	// Checking if member exist in guild
	member, err = s.GetMember(context.Background(), guildID, disgord.NewSnowflake(id))
	if err != nil {
		return &disgord.Member{}, false
	}

	return member, true
}
