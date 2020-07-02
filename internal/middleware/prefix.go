package middleware

import (
	"context"

	"github.com/Terisback/robo-biba/internal/command"
	"github.com/andersfylling/disgord"
)

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
