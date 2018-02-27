package main

import (
	"context"
	"os"

	"github.com/andviro/middleware"
	"github.com/go-mixins/log/logrus"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	tg "github.com/go-mixins/bot/telegram"
)

func main() {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		logger.Fatalf("please set environment variable BOT_TOKEN")
	}
	b, err := tg.New(token)
	if err != nil {
		logger.Fatalf("%+v", err)
	}
	b.Concurrency = 100
	b.Use(func(ctx context.Context, next middleware.Handler) (err error) {
		return next.Apply(ctx)
	})
	b.On(tg.Command("start"), func(ctx context.Context) error {
		return b.Reply(ctx, "Hello!")
	})
	b.On(tg.Command("quit"), func(ctx context.Context) error {
		return b.Reply(ctx, "Bye!")
	})
	b.On(tg.Command("list"), func(ctx context.Context) error {
		return b.Reply(ctx, "sample keyboard", func(msg *tgbotapi.MessageConfig) {
			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("1", "1"),
					tgbotapi.NewInlineKeyboardButtonData("2", "2"),
					tgbotapi.NewInlineKeyboardButtonData("3", "3"),
				),
			)
		})
	})
	b.On(tg.Action("1"), func(ctx context.Context) error {
		return b.EditMessageText(ctx, "sample keyboard (1)", func(msg *tgbotapi.EditMessageTextConfig) {
			markup := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("1", "1"),
					tgbotapi.NewInlineKeyboardButtonData("2", "2"),
					tgbotapi.NewInlineKeyboardButtonData("3", "3"),
				),
			)
			msg.ReplyMarkup = &markup
		})
	})
	b.On(tg.Action("2"), func(ctx context.Context) error {
		return b.EditMessageText(ctx, "sample keyboard (2)")
	})
	b.On(tg.Action("3"), func(ctx context.Context) error {
		return b.EditMessageText(ctx, "sample keyboard (1)", func(msg *tgbotapi.EditMessageTextConfig) {
			markup := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("1", "1"),
					tgbotapi.NewInlineKeyboardButtonData("2", "2"),
				),
			)
			msg.ReplyMarkup = &markup
		})
	})
	b.On(tg.Hears("lol"), func(ctx context.Context) error {
		return b.Reply(ctx, "lol yourself", b.WithReply(ctx))
	})
	b.On(tg.Update("Text"), func(ctx context.Context) error {
		return b.Reply(ctx, b.Message(ctx).Text)
	})
	b.On(tg.Update("Sticker"), func(ctx context.Context) error {
		return b.Reply(ctx, "sticker!")
	})
	b.On(tg.Update("NewChatMembers"), func(ctx context.Context) error {
		for _, u := range *b.Message(ctx).NewChatMembers {
			if u.UserName == b.Self.UserName {
				return b.Reply(ctx, "Hello all!")
			}
			if err := b.Reply(ctx, "Hi, "+u.UserName+"!"); err != nil {
				return err
			}
		}
		return nil
	})
	b.On(tg.Update("LeftChatMember"), func(ctx context.Context) error {
		u := b.Message(ctx).LeftChatMember
		if u.UserName == b.Self.UserName {
			logger.Warnf("Was kicked from %s by %s", b.Chat(ctx).Title, b.From(ctx))
			return nil
		}
		return b.Reply(ctx, "Bye, "+u.UserName+"!")
	})
	if err = b.Run(); err != nil {
		logger.Fatalf("%+v", err)
	}
}
