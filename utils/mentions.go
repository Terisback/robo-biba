package utils

import (
	"context"
	"strconv"

	"github.com/andersfylling/disgord"
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

	if UserIDRegex.MatchString(arg) {
		userID := UserIDRegex.FindString(arg)
		id, err = strconv.ParseUint(userID, 10, 64)
		if err != nil {
			return 0, false
		}
		return id, true
	}

	return 0, false
}

func GetIDFromArgAndCheckIt(s disgord.Session, guildID disgord.Snowflake, arg string) (id uint64, ok bool) {
	id, ok = GetIDFromArg(arg)
	if !ok {
		return
	}

	// Checking if member exist in guild
	_, err := s.GetMember(context.Background(), guildID, disgord.NewSnowflake(id))
	if err != nil {
		return id, false
	}

	return id, true
}
