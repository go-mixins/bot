package telegram

import (
	"context"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/go-mixins/bot"
)

func (drv *Driver) Reply(ctx context.Context, obj interface{}) (err error) {
	upd, ok := ctx.Value(botKey).(tgbotapi.Update)
	if !ok {
		return bot.Errors.New("no update in context")
	}
	var message *tgbotapi.Message
	switch {
	case upd.Message != nil:
		message = upd.Message
	default:
		return bot.Errors.New("no message in context")
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
