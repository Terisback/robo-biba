package middleware

import (
	"context"

	"github.com/Terisback/robo-biba/utils"
	"github.com/andersfylling/disgord"
)

// FilterPrefix checks if a message begins with a prefix
func (h *Middleware) FilterPrefix(prefixes ...string) func(evt interface{}) interface{} {
	return func(evt interface{}) interface{} {
		// Check the needed event type
		if e, ok := evt.(*disgord.MessageCreate); ok {
			content := e.Message.Content
			selfID := h.self.ID.String()

			// Parsing arguments (fields) from the message
			args, ok := utils.ParseArguments(content, selfID, prefixes...)
			if ok {
				// Setting arguments in the context
				e.Ctx = context.WithValue(e.Ctx, utils.ArgumentsContextKey, args)
				return evt
			}
		}

		// Break the route
		return nil
	}
}
