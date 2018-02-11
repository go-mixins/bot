package main

import (
	"context"
	"os"

	"github.com/andviro/middleware"
	"github.com/go-mixins/log"
	"github.com/go-mixins/log/logrus"

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
	b.Use(func(ctx context.Context, next middleware.Handler) (err error) {
		msg := b.Message(ctx)
		if msg != nil {
			logger.WithContext(log.M{
				"from": msg.From,
			}).Debug(msg.Text)
		}
		return next.Apply(ctx)
	})
	b.On(b.Command("start"), func(ctx context.Context) error {
		return b.Reply(ctx, "Hello!")
	})
	b.On(b.Command("quit"), func(ctx context.Context) error {
		return b.Reply(ctx, "Bye!")
	})
	b.On(b.Hears("lol"), func(ctx context.Context) error {
		return b.Reply(ctx, "lol yourself", b.WithReply)
	})
	b.On(b.Update("Text"), func(ctx context.Context) error {
		return b.Reply(ctx, b.Message(ctx).Text)
	})
	b.On(b.Update("Sticker"), func(ctx context.Context) error {
		return b.Reply(ctx, "sticker!")
	})
	b.On(b.Update("NewChatMembers"), func(ctx context.Context) error {
		for _, u := range *b.Message(ctx).NewChatMembers {
			if u.UserName == b.Self.UserName {
				continue
			}
			if err := b.Reply(ctx, "Hi, "+u.UserName+"!"); err != nil {
				return err
			}
		}
		return nil
	})
	b.On(b.Update("LeftChatMember"), func(ctx context.Context) error {
		u := b.Message(ctx).LeftChatMember
		return b.Reply(ctx, "Bye, "+u.UserName+"!")
	})
	if err = b.Run(); err != nil {
		logger.Fatalf("%+v", err)
	}
}
