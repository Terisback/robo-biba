package middleware

import (
	"context"
	"errors"

	"github.com/andersfylling/disgord"
)

type Holder struct {
	session disgord.Session
	self    *disgord.User
}

// Create new middleware holder
func New(s disgord.Session) (*Holder, error) {
	var (
		user *disgord.User
		err  error
	)
	if user, err = s.GetCurrentUser(context.Background()); err != nil {
		return nil, errors.New("Unable to fetch info about the bot instance")
	}
	return &Holder{session: s, self: user}, nil
}
