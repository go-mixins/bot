package telegram

import (
	"context"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/go-mixins/bot"
)

func (drv *Driver) Reply(ctx context.Context, text string, opts ...bot.Option) (err error) {
	chat := drv.Chat(ctx)
	if chat == nil {
		return bot.Errors.New("no chat to reply in")
	}
	msg := tgbotapi.NewMessage(chat.ID, text)
	for _, opt := range opts {
		opt(ctx, &msg)
	}
	_, err = drv.Send(msg)
	if err != nil {
		err = bot.Errors.Wrap(err, "sending message")
	}
	return
}
