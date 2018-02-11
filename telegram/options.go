package telegram

import "context"

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (drv *Driver) WithReply(ctx context.Context, args ...interface{}) {
	if len(args) == 0 {
		return
	}
	msg := drv.message(ctx)
	if msg == nil {
		return
	}
	dest, ok := args[0].(*tgbotapi.MessageConfig)
	if !ok {
		return
	}
	dest.ReplyToMessageID = msg.MessageID
}
