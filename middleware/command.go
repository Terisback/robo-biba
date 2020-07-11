package middleware

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

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

			// Parsing command arguments
			cmd, err := command(content, selfID, commandOptions)
			if err != nil {
				return nil
			}

			// Passing command arguments to context
			e.Ctx = context.WithValue(e.Ctx, CommandKey, cmd)

			return evt
		}

		// Break the route
		return nil
	}
}

type Argument struct {
	value   string
	numeric *int
	id      *uint64
}

func (a *Argument) String() string {
	return a.value
}

func (a *Argument) GetNumeric() (value int, ok bool) {
	if a.numeric == nil {
		return 0, false
	}

	return *a.numeric, true
}

func (a *Argument) GetID() (id uint64, ok bool) {
	if a.id == nil {
		return 0, false
	}

	return *a.id, true
}

type Command struct {
	Arguments []Argument
}

var (
	numericRegex = regexp.MustCompile(`\d+`)
)

func command(content, selfID string, options CommandOptions) (command Command, err error) {
	// Normalize
	content = strings.ToLower(content)
	content = strings.TrimLeft(content, " ")
	selfMention := fmt.Sprintf(`<@%s>`, selfID)

	// Find the prefixes and trim them
	if prefix, ok := hasOneOfThePrefixes(content, options.Prefixes...); ok {
		// Usual prefix
		content = strings.TrimPrefix(content, prefix)
	} else if strings.HasPrefix(content, selfMention) {
		// Mention prefix
		content = strings.TrimPrefix(content, selfMention)
	} else {
		// There is no prefix
		return command, errors.New("There is no any of the prefixes")
	}

	if ok := utils.Aliases(content, options.Aliases...); !ok {
		return command, errors.New("There is no any of the aliases")
	}

	// Command fields
	fields := strings.Fields(content)

	for _, f := range fields {
		arg := Argument{}
		arg.value = f

		if id, ok := utils.GetIDFromArg(f); ok {
			arg.id = &id
		} else if numericRegex.MatchString(f) {
			numStr := numericRegex.FindString(f)
			num, err := strconv.ParseInt(numStr, 10, 64)
			if err != nil {
				command.Arguments = append(command.Arguments, arg)
				continue
			}
			n := int(num)
			arg.numeric = &n
		}

		command.Arguments = append(command.Arguments, arg)
	}

	return command, nil
}

func hasOneOfThePrefixes(content string, prefixes ...string) (prefix string, ok bool) {
	for _, prefix = range prefixes {
		prefix = strings.ToLower(prefix)
		prefix = strings.TrimSpace(prefix)
		if ok = strings.HasPrefix(content, prefix); ok {
			return
		}
	}

	return "", false
}

type Key int

const (
	CommandKey Key = 0
)

// Allows you to get command from context after FilterCommand
func GetCommandFromContext(ctx context.Context) (Command, error) {
	cmd := ctx.Value(CommandKey)

	if cmd == nil {
		return Command{}, errors.New("Command from context is nil")
	}

	return cmd.(Command), nil
}
