package telegram

import (
	"context"
	"regexp"
	"strings"

	"github.com/andviro/middleware"
	"github.com/fatih/structs"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func Command(cmd ...string) middleware.Predicate {
	cmdRe := regexp.MustCompile("^" + strings.Join(cmd, "|") + "$")
	return func(ctx context.Context) bool {
		upd, _ := ctx.Value(botKey).(tgbotapi.Update)
		if upd.Message != nil {
			return upd.Message.IsCommand() && cmdRe.MatchString(upd.Message.Command())
		}
		return false
	}
}

func Hears(word string) middleware.Predicate {
	wordRe := regexp.MustCompile(word)
	return func(ctx context.Context) bool {
		upd, _ := ctx.Value(botKey).(tgbotapi.Update)
		switch {
		case upd.Message != nil:
			return wordRe.MatchString(upd.Message.Caption) || wordRe.MatchString(upd.Message.Text)
		case upd.CallbackQuery != nil:
			return wordRe.MatchString(upd.CallbackQuery.Data)
		}
		return false
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

func Action(q string) middleware.Predicate {
	return func(ctx context.Context) bool {
		upd, _ := ctx.Value(botKey).(tgbotapi.Update)
		if upd.CallbackQuery == nil {
			return false
		}
		return upd.CallbackQuery.Data == q
	}
}
