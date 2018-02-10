package telegram

import (
	"context"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/go-mixins/bot"
)

func (drv *Driver) Reply(ctx context.Context, obj interface{}) (err error) {
	msg := drv.message(ctx)
	if msg == nil || msg.Chat == nil {
		return bot.Errors.New("no chat to reply in")
	}
	msgID := drv.msgID(ctx)
	var v tgbotapi.Chattable
	switch t := obj.(type) {
	case string:
		msg := tgbotapi.NewMessage(msg.Chat.ID, t)
		if msgID != 0 {
			msg.ReplyToMessageID = msgID
		}
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
