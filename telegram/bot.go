package telegram

import (
	"context"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/go-mixins/bot"
	"github.com/go-mixins/bot/generic"
)

type botKeyType int

const (
	botKey botKeyType = iota
	meKey
)

type Driver struct {
	*tgbotapi.BotAPI
	stop    chan struct{}
	updates tgbotapi.UpdatesChannel
}

func (drv *Driver) Close() {
	close(drv.stop)
}

func (drv *Driver) Next() bool {
	select {
	case <-drv.stop:
		return false
	default:
		return true
	}
}

func (drv *Driver) Context() context.Context {
	return context.WithValue(context.Background(), botKey, <-drv.updates)
}

var _ bot.Driver = (*Driver)(nil)

type Bot struct {
	Driver
	generic.Bot
}

var _ bot.Bot = (*Bot)(nil)

func New(token string) (res *Bot, err error) {
	res = new(Bot)
	res.Driver.stop = make(chan struct{})
	if res.Driver.BotAPI, err = tgbotapi.NewBotAPI(token); err != nil {
		err = bot.Errors.Wrap(err, "creating telegram bot")
		return
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	if res.Driver.updates, err = res.Driver.GetUpdatesChan(u); err != nil {
		err = bot.Errors.Wrap(err, "getting updates")
		return
	}
	return
}

func (b *Bot) Run() error {
	return b.Bot.Run(&b.Driver)
}
