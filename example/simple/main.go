package main

import (
	"context"
	"os"

	"github.com/go-mixins/log"
	"github.com/go-mixins/log/logrus"

	"github.com/go-mixins/bot/telegram"
)

func main() {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		logger.Fatalf("please set environment variable BOT_TOKEN")
	}
	bot, err := telegram.New(token)
	if err != nil {
		logger.Fatalf("%+v", err)
	}
	defer bot.Close()
	bot.On(bot.Command("/start"), func(ctx context.Context) error {
		return bot.Reply(ctx, "Hello!")
	})
	bot.On(bot.Command("/quit"), func(ctx context.Context) error {
		return bot.Reply(ctx, "Bye!")
	})
	bot.On(bot.Message, func(ctx context.Context) error {
		from := bot.From(ctx)
		if from != nil {
			logger.WithContext(log.M{
				"from": from.UserName,
			}).Infof("message: %s", bot.Text(ctx))
			return nil
		}
		logger.Warnf("anonymous message: ", bot.Text(ctx))
		return nil
	})
	if err = bot.Run(); err != nil {
		logger.Fatalf("%+v", err)
	}
}
