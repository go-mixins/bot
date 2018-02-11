package telegram

import (
	"context"
	"encoding/json"

	"github.com/go-mixins/bot"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (drv *Driver) Msg(ctx context.Context) *bot.Message {
	msg := drv.message(ctx)
	if msg == nil {
		return nil
	}
	return &bot.Message{
		Text: msg.Text,
		Meta: bot.Meta{
			Platform: "telegram",
		},
	}
}

func (drv *Driver) user(u *tgbotapi.User) *bot.User {
	if u == nil {
		return nil
	}
	return &bot.User{
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		UserName:     u.UserName,
		LanguageCode: u.LanguageCode,
		Meta: bot.Meta{
			Platform: "telegram",
		},
	}
}

func (drv *Driver) chat(ctx context.Context) (res *tgbotapi.Chat) {
	upd, _ := ctx.Value(botKey).(tgbotapi.Update)
	switch {
	case upd.Message != nil:
		res = upd.Message.Chat
	case upd.CallbackQuery != nil && upd.CallbackQuery.Message != nil:
		res = upd.CallbackQuery.Message.Chat
	case upd.ChannelPost != nil:
		res = upd.ChannelPost.Chat
	case upd.EditedChannelPost != nil:
		res = upd.EditedChannelPost.Chat
	case upd.EditedMessage != nil:
		res = upd.EditedMessage.Chat
	}
	return
}

func (drv *Driver) Chat(ctx context.Context) *bot.Chat {
	chat := drv.chat(ctx)
	if chat == nil {
		return nil
	}
	return &bot.Chat{
		Meta: bot.Meta{
			Platform: "telegram",
		},
		Type:        chat.Type,
		Title:       chat.Title,
		UserName:    chat.UserName,
		FirstName:   chat.FirstName,
		LastName:    chat.LastName,
		Description: chat.Description,
	}
}

func (drv *Driver) message(ctx context.Context) (res *tgbotapi.Message) {
	upd, _ := ctx.Value(botKey).(tgbotapi.Update)
	return upd.Message
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

func (drv *Driver) Me(ctx context.Context) *bot.User {
	return drv.user(&drv.api.Self)
}

func (drv *Driver) From(ctx context.Context) *bot.User {
	return drv.user(drv.from(ctx))
}

func (drv *Driver) Debug(ctx context.Context) string {
	upd, _ := ctx.Value(botKey).(tgbotapi.Update)
	jsonData, _ := json.MarshalIndent(upd, "", "  ")
	return string(jsonData)
}
