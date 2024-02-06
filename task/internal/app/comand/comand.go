package command

import (
	"context"
	"fmt"
	"strings"
)

type CommandHandler[c any] interface {
	Handle(ctx context.Context, cmd c) error
}

func getCommandName(cmd any) string {
	return strings.Split(fmt.Sprintf("%T", cmd), ".")[1]
}
