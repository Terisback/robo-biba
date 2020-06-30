package middleware

import (
	"context"
	"errors"
	"strings"

	"github.com/Terisback/robo-biba/internal/command"
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

// FilterBotMessage filtering bot messages
func (h *Holder) FilterBotMessages(evt interface{}) interface{} {
	if e, ok := evt.(*disgord.MessageCreate); ok {
		if e.Message.Author.Bot {
			return nil
		}
	}
	return evt
}

// FilterPrefix bypassing by prefix, and passing to context arguments of command (need for FilterAliases)
func (h *Holder) FilterPrefix(prefix string) func(evt interface{}) interface{} {
	return func(evt interface{}) interface{} {
		if e, ok := evt.(*disgord.MessageCreate); ok {
			content := e.Message.Content

			args, ok := command.Parse(h.self.ID.String(), prefix, content)
			if ok {
				e.Ctx = context.WithValue(e.Ctx, "args", args)
				return evt
			}
		}

		return nil
	}
}

// FilterAliases bypassing by aliases, AND WORKS ONLY AFTER FilterPrefix
func (h *Holder) FilterAliases(aliases ...string) func(evt interface{}) interface{} {
	return func(evt interface{}) interface{} {
		if e, ok := evt.(*disgord.MessageCreate); ok {
			// Getting args that passed in FilterPrefix
			args := e.Ctx.Value("args").(command.Arguments)

			if args == nil || len(args) == 0 {
				return nil
			}

			command := args[0]

			for _, alias := range aliases {
				// Normalizing alias
				alias = strings.TrimSpace(alias)
				alias = strings.ToLower(alias)

				// If command equal to it alias; pass to next middleware\handler
				if command == alias {
					return evt
				}
			}
		}

		return nil
	}
}
