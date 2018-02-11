package telegram

import "context"

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (drv *Driver) WithReply(ctx context.Context, arg interface{}) {
	dest, ok := arg.(*tgbotapi.MessageConfig)
	if !ok {
		return
	}
	if msg := drv.Message(ctx); msg != nil {
		dest.ReplyToMessageID = msg.MessageID
	}
}
