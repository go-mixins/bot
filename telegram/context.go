package telegram

import (
	"context"
	"encoding/json"

	"github.com/go-mixins/bot"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (drv *Driver) Text(ctx context.Context) (res string) {
	msg := drv.message(ctx)
	if msg == nil {
		return
	}
	if msg.Caption != "" {
		return msg.Caption
	}
	return msg.Text
}

func (drv *Driver) message(ctx context.Context) (res *tgbotapi.Message) {
	upd, _ := ctx.Value(botKey).(tgbotapi.Update)
	switch {
	case upd.Message != nil:
		res = upd.Message
	case upd.CallbackQuery != nil:
		res = upd.CallbackQuery.Message
	case upd.ChannelPost != nil:
		res = upd.ChannelPost
	case upd.EditedChannelPost != nil:
		res = upd.EditedChannelPost
	case upd.EditedMessage != nil:
		res = upd.EditedMessage
	}
	return
}

func (drv *Driver) from(ctx context.Context) (from *tgbotapi.User) {
	upd, _ := ctx.Value(botKey).(tgbotapi.Update)
	switch {
	case upd.Message != nil:
		from = upd.Message.From
	case upd.CallbackQuery != nil:
		from = upd.CallbackQuery.From
	case upd.ChannelPost != nil:
		from = upd.ChannelPost.From
	case upd.ChosenInlineResult != nil:
		from = upd.ChosenInlineResult.From
	case upd.EditedChannelPost != nil:
		from = upd.EditedChannelPost.From
	case upd.EditedMessage != nil:
		from = upd.EditedMessage.From
	case upd.InlineQuery != nil:
		from = upd.InlineQuery.From
	case upd.PreCheckoutQuery != nil:
		from = upd.PreCheckoutQuery.From
	case upd.ShippingQuery != nil:
		from = upd.ShippingQuery.From
	}
	return
}

func (drv *Driver) From(ctx context.Context) *bot.User {
	from := drv.from(ctx)
	if from == nil {
		return nil
	}
	return &bot.User{
		FirstName:    from.FirstName,
		LastName:     from.LastName,
		UserName:     from.UserName,
		LanguageCode: from.LanguageCode,
	}
}

func (drv *Driver) Debug(ctx context.Context) string {
	upd, _ := ctx.Value(botKey).(tgbotapi.Update)
	jsonData, _ := json.MarshalIndent(upd, "", "  ")
	return string(jsonData)
}
