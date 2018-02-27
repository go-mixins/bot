package telegram

import "context"

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (drv *Driver) WithReply(ctx context.Context) func(*tgbotapi.MessageConfig) {
	return func(dest *tgbotapi.MessageConfig) {
		if msg := drv.Message(ctx); msg != nil {
			dest.ReplyToMessageID = msg.MessageID
		}
	}
}
