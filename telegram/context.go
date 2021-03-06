package telegram

import (
	"context"
	"encoding/json"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (Driver) From(ctx context.Context) (from *tgbotapi.User) {
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

func (Driver) Chat(ctx context.Context) (res *tgbotapi.Chat) {
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

func (Driver) Command(ctx context.Context) (res string) {
	upd, _ := ctx.Value(botKey).(tgbotapi.Update)
	if upd.Message != nil {
		return upd.Message.Command()
	}
	return
}

func (Driver) Arguments(ctx context.Context) (res string) {
	upd, _ := ctx.Value(botKey).(tgbotapi.Update)
	switch {
	case upd.Message != nil:
		msg := upd.Message
		if msg.IsCommand() {
			return msg.CommandArguments()
		}
		return msg.Text
	case upd.CallbackQuery != nil:
		return upd.CallbackQuery.Data
	}
	return
}

func (Driver) Message(ctx context.Context) (res *tgbotapi.Message) {
	upd, _ := ctx.Value(botKey).(tgbotapi.Update)
	return upd.Message
}

func (Driver) Update(ctx context.Context) (res tgbotapi.Update) {
	upd, _ := ctx.Value(botKey).(tgbotapi.Update)
	return upd
}

func (Driver) SaveContext(ctx context.Context) ([]byte, error) {
	upd, _ := ctx.Value(botKey).(tgbotapi.Update)
	return json.Marshal(upd)
}

func (Driver) RestoreContext(ctx context.Context, data []byte) (res context.Context, err error) {
	var upd tgbotapi.Update
	if err = json.Unmarshal(data, &upd); err != nil {
		return
	}
	res = context.WithValue(ctx, botKey, upd)
	return
}
