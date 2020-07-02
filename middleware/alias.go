package middleware

import (
	"log"

	"github.com/Terisback/robo-biba/command"
	"github.com/andersfylling/disgord"
)

// FilterAliases bypassing by aliases, AND WORKS ONLY AFTER FilterPrefix
func (h *Holder) FilterAliases(aliases ...string) func(evt interface{}) interface{} {
	return func(evt interface{}) interface{} {
		if e, ok := evt.(*disgord.MessageCreate); ok {
			// Getting args that passed in FilterPrefix
			args, err := command.GetArgsFromContext(e.Ctx)
			if err != nil {
				log.Println(err)
				return nil
			}

			// First argument would be a command itself
			com := args[0]

			if command.Aliases(com, aliases...) {
				return evt
			}
		}

		return nil
	}
}
