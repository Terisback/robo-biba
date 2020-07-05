package middleware

import (
	"context"

	"github.com/Terisback/robo-biba/utils"
	"github.com/andersfylling/disgord"
)

type CommandOptions struct {
	Prefixes []string
	Aliases  []string
}

// FilterCommand checks if a message begins with a prefix and if one of the command aliases exists after the prefix
func (h *Middleware) FilterCommand(commandOptions CommandOptions) func(evt interface{}) interface{} {
	return func(evt interface{}) interface{} {
		// Check the needed event type
		if e, ok := evt.(*disgord.MessageCreate); ok {
			content := e.Message.Content
			selfID := h.self.ID.String()

			// Parsing arguments (fields) from the message
			args, ok := utils.ParseArguments(content, selfID, commandOptions.Prefixes...)
			if !ok || len(args) == 0 {
				// Break the route
				return nil
			}

			// First argument would be a command itself
			command := args[0]

			// Checking command aliases
			if utils.Aliases(command, commandOptions.Aliases...) {
				// Setting arguments in the context
				e.Ctx = context.WithValue(e.Ctx, utils.ArgumentsContextKey, args)

				// Continue to execute the route
				return evt
			}
		}

		// Break the route
		return nil
	}
}
