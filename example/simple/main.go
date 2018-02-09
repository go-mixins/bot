package main

import (
	"context"
	"os"

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
	if err = bot.Run(); err != nil {
		logger.Fatalf("%+v", err)
	}
}
