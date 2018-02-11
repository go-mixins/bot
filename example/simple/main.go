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
	b.Concurrency = 100
	b.Use(func(ctx context.Context, next middleware.Handler) (err error) {
		msg := tg.Message(ctx)
		if msg != nil {
			logger.WithContext(log.M{
				"from": msg.From,
			}).Debug(msg.Text)
		}
		return next.Apply(ctx)
	})
	b.On(tg.Command("start"), func(ctx context.Context) error {
		return b.Reply(ctx, "Hello!")
	})
	b.On(tg.Command("quit"), func(ctx context.Context) error {
		return b.Reply(ctx, "Bye!")
	})
	b.On(tg.Hears("lol"), func(ctx context.Context) error {
		return b.Reply(ctx, "lol yourself", b.WithReply)
	})
	b.On(tg.Update("Text"), func(ctx context.Context) error {
		return b.Reply(ctx, tg.Message(ctx).Text)
	})
	b.On(tg.Update("Sticker"), func(ctx context.Context) error {
		return b.Reply(ctx, "sticker!")
	})
	b.On(tg.Update("NewChatMembers"), func(ctx context.Context) error {
		for _, u := range *tg.Message(ctx).NewChatMembers {
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
		u := tg.Message(ctx).LeftChatMember
		if u.UserName == b.Self.UserName {
			logger.Warnf("Was kicked from %s by %s", tg.Chat(ctx).Title, tg.From(ctx))
			return nil
		}
		return b.Reply(ctx, "Bye, "+u.UserName+"!")
	})
	if err = b.Run(); err != nil {
		logger.Fatalf("%+v", err)
	}
}
