package middleware

import (
	"context"
	"errors"

	"github.com/andersfylling/disgord"
)

type Middleware struct {
	session disgord.Session
	self    *disgord.User
}

// Create new Middleware
func New(session disgord.Session) (*Middleware, error) {
	var (
		user *disgord.User
		err  error
	)
	if user, err = session.GetCurrentUser(context.Background()); err != nil {
		return nil, errors.New("Unable to fetch info about the bot instance")
	}
	return &Middleware{session: session, self: user}, nil
}
