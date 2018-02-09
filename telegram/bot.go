package telegram

import (
	"context"

	"github.com/andviro/middleware"
	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/go-mixins/bot"
	"github.com/go-mixins/bot/generic"
)

type botKeyType int

const (
	botKey botKeyType = iota
)

type Driver struct {
	api     *tgbotapi.BotAPI
	stop    chan struct{}
	closed  bool
	updates tgbotapi.UpdatesChannel
}

var _ bot.Driver = (*Driver)(nil)

func New(token string) (res *generic.Bot, err error) {
	driver := &Driver{
		stop: make(chan struct{}),
	}
	driver.api, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		err = bot.Errors.Wrap(err, "creating telegram bot")
		return
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	driver.updates, err = driver.api.GetUpdatesChan(u)
	if err != nil {
		err = bot.Errors.Wrap(err, "getting updates")
		return
	}
	return generic.New(driver)
}

func (drv *Driver) Close() error {
	close(drv.stop)
	return nil
}

func (drv *Driver) Context(ctx context.Context, next middleware.Handler) error {
	return next.Apply(context.WithValue(ctx, botKey, <-drv.updates))
}

func (drv *Driver) Next() bool {
	select {
	case <-drv.stop:
		return false
	default:
		return true
	}
}
