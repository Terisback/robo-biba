package middleware

import "github.com/andersfylling/disgord"

// FilterBotMessage doesn't bypass on messages from bots
func (h *Middleware) FilterBotMessages(evt interface{}) interface{} {
	if e, ok := evt.(*disgord.MessageCreate); ok {
		if e.Message.Author.Bot {
			return nil
		}
	}
	return evt
}
