package telegram

import (
	"context"

	"github.com/andviro/middleware"
	"github.com/fatih/structs"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func Command(cmd string) middleware.Predicate {
	return func(ctx context.Context) bool {
		if msg := Message(ctx); msg != nil {
			return msg.IsCommand() && cmd == msg.Command()
		}
		return false
	}
}

func Hears(word string) middleware.Predicate {
	return func(ctx context.Context) bool {
		upd, _ := ctx.Value(botKey).(tgbotapi.Update)
		if upd.Message != nil {
			if upd.Message.Caption == word {
				return true
			}
			if upd.Message.Text == word {
				return true
			}
		}
		return upd.CallbackQuery != nil && upd.CallbackQuery.Data == word
	}
}

func Update(field string) middleware.Predicate {
	return func(ctx context.Context) bool {
		upd, _ := ctx.Value(botKey).(tgbotapi.Update)
		if f, ok := structs.New(&upd).FieldOk(field); ok {
			return !f.IsZero()
		}
		if upd.Message != nil {
			if f, ok := structs.New(upd.Message).FieldOk(field); ok {
				return !f.IsZero()
			}
		}
		return false
	}
}
