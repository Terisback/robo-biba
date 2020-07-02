package middleware

import "github.com/andersfylling/disgord"

// FilterBotMessage filtering bot messages
func (h *Holder) FilterBotMessages(evt interface{}) interface{} {
	if e, ok := evt.(*disgord.MessageCreate); ok {
		if e.Message.Author.Bot {
			return nil
		}
	}
	return evt
}
