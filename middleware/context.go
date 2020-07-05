package middleware

import (
	"context"
	"errors"

	"github.com/Terisback/robo-biba/utils"
)

// Allows you to get arguments from context after FilterCommand
// Return's nil args if it's zero-length
func GetArgsFromContext(ctx context.Context) (utils.Arguments, error) {
	args := ctx.Value(utils.ArgumentsContextKey).(utils.Arguments)

	if args == nil {
		return nil, errors.New("Arguments from context is nil")
	}

	return args, nil
}
