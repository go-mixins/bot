package telegram

import (
	"context"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/go-mixins/bot"
)

func (b *Bot) Reply(ctx context.Context, text string, opts ...bot.MessageOption) (err error) {
	chat := b.Chat(ctx)
	if chat == nil {
		return bot.Errors.New("no chat to reply in")
	}
	msg := tgbotapi.NewMessage(chat.ID, text)
	for _, opt := range opts {
		opt(&msg)
	}
	_, err = b.Send(msg)
	err = bot.Errors.Wrap(err, "sending message")
	return
}

func (b *Bot) EditMessageText(ctx context.Context, text string, opts ...bot.EditMessageOption) (err error) {
	upd, _ := ctx.Value(botKey).(tgbotapi.Update)
	if upd.CallbackQuery == nil || upd.CallbackQuery.Message == nil || upd.CallbackQuery.Message.Chat == nil {
		return nil
	}
	editMsg := tgbotapi.NewEditMessageText(upd.CallbackQuery.Message.Chat.ID, upd.CallbackQuery.Message.MessageID, text)
	for _, opt := range opts {
		opt(&editMsg)
	}
	_, err = b.Send(editMsg)
	err = bot.Errors.Wrap(err, "sending message edit")
	return
}
