package telegram

import (
	"context"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (drv *Driver) Text(ctx context.Context) string {
	upd, _ := ctx.Value(botKey).(tgbotapi.Update)
	switch {
	case upd.Message != nil:
		return upd.Message.Text
	case upd.ChannelPost != nil:
		return upd.ChannelPost.Text
	}
	return ""
}

func (drv *Driver) UserName(ctx context.Context) string {
	upd, _ := ctx.Value(botKey).(tgbotapi.Update)
	var from *tgbotapi.User
	switch {
	case upd.Message != nil:
		from = upd.Message.From
	case upd.ChannelPost != nil:
		from = upd.ChannelPost.From
	}
	if from != nil {
		return from.UserName
	}
	return ""
}
