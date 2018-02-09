package telegram

import (
	"context"
	"strings"

	"github.com/andviro/middleware"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (drv *Driver) Command(cmd string) middleware.Predicate {
	return func(ctx context.Context) bool {
		upd, ok := ctx.Value(botKey).(tgbotapi.Update)
		if !ok {
			return false
		}
		var text string
		switch {
		case upd.Message != nil:
			text = upd.Message.Text
		}
		text = strings.TrimSpace(text)
		if len(text) == 0 || !strings.HasPrefix(text, "/") {
			return false
		}
		args := strings.SplitN(text, " ", 2)
		return args[0] == cmd
	}
}

func (drv *Driver) Hears(word string) middleware.Predicate {
	return func(ctx context.Context) bool {
		upd, ok := ctx.Value(botKey).(tgbotapi.Update)
		if !ok {
			return false
		}
		var text string
		switch {
		case upd.Message != nil:
			text = upd.Message.Text
		}
		return strings.Contains(text, word)
	}
}

func (drv *Driver) Message(ctx context.Context) bool {
	upd, ok := ctx.Value(botKey).(tgbotapi.Update)
	if !ok {
		return false
	}
	return upd.Message != nil
}
