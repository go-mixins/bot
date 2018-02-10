package telegram

import (
	"context"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/go-mixins/bot"
)

func (drv *Driver) Reply(ctx context.Context, obj interface{}) (err error) {
	message := drv.message(ctx)
	if message == nil || message.Chat == nil {
		return bot.Errors.New("no chat to reply in")
	}
	var v tgbotapi.Chattable
	switch t := obj.(type) {
	case string:
		msg := tgbotapi.NewMessage(message.Chat.ID, t)
		msg.ReplyToMessageID = message.MessageID
		v = msg
	default:
		return bot.Errors.New("unsupported send object")
	}
	_, err = drv.api.Send(v)
	if err != nil {
		err = bot.Errors.Wrap(err, "sending message")
	}
	return
}
