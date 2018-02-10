package telegram

import (
	"context"

	"github.com/andviro/middleware"
	"github.com/go-mixins/bot"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (drv *Driver) Command(cmd string) middleware.Predicate {
	return func(ctx context.Context) bool {
		msg := drv.message(ctx)
		if msg == nil || !msg.IsCommand() {
			return false
		}
		return msg.Command() == cmd
	}
}

func (drv *Driver) Hears(word string) middleware.Predicate {
	return func(ctx context.Context) bool {
		return drv.Text(ctx) == word
	}
}

func (drv *Driver) Message(t bot.MessageType) middleware.Predicate {
	return func(ctx context.Context) bool {
		msg := drv.message(ctx)
		if msg == nil {
			return false
		}
		switch t {
		case bot.MsgVoice:
			return msg.Voice != nil
		case bot.MsgVideoNote:
			return msg.VideoNote != nil
		case bot.MsgVideo:
			return msg.Video != nil
		case bot.MsgVenue:
			return msg.Venue != nil
		case bot.MsgText:
			return msg.Text != ""
		case bot.MsgSupergroupChatCreated:
			return msg.SuperGroupChatCreated
		case bot.MsgSuccessfulPayment:
			return msg.SuccessfulPayment != nil
		case bot.MsgSticker:
			return msg.Sticker != nil
		case bot.MsgPinnedMessage:
			return msg.PinnedMessage != nil
		case bot.MsgPhoto:
			return msg.Photo != nil
		case bot.MsgNewChatTitle:
			return msg.NewChatTitle != ""
		case bot.MsgNewChatPhoto:
			return msg.NewChatPhoto != nil
		case bot.MsgNewChatMembers:
			return msg.NewChatMembers != nil
		case bot.MsgMigrateToChatID:
			return msg.MigrateToChatID != 0
		case bot.MsgMigrateFromChatID:
			return msg.MigrateFromChatID != 0
		case bot.MsgLocation:
			return msg.Location != nil
		case bot.MsgLeftChatMember:
			return msg.LeftChatMember != nil
		case bot.MsgInvoice:
			return msg.Invoice != nil
		case bot.MsgGroupChatCreated:
			return msg.GroupChatCreated
		case bot.MsgGame:
			return msg.Game != nil
		case bot.MsgDocument:
			return msg.Document != nil
		case bot.MsgDeleteChatPhoto:
			return msg.DeleteChatPhoto
		case bot.MsgContact:
			return msg.Contact != nil
		case bot.MsgChannelChatCreated:
			return msg.ChannelChatCreated
		case bot.MsgAudio:
			return msg.Audio != nil
		}
		return false
	}
}

func (drv *Driver) Update(t bot.UpdateType) middleware.Predicate {
	return func(ctx context.Context) bool {
		upd, _ := ctx.Value(botKey).(tgbotapi.Update)
		switch t {
		case bot.UpdCallbackQuery:
			return upd.CallbackQuery != nil
		case bot.UpdChannelPost:
			return upd.ChannelPost != nil
		case bot.UpdChosenInlineResult:
			return upd.ChosenInlineResult != nil
		case bot.UpdEditedChannelPost:
			return upd.EditedChannelPost != nil
		case bot.UpdEditedMessage:
			return upd.EditedMessage != nil
		case bot.UpdInlineQuery:
			return upd.InlineQuery != nil
		case bot.UpdShippingQuery:
			return upd.ShippingQuery != nil
		case bot.UpdPreCheckoutQuery:
			return upd.PreCheckoutQuery != nil
		case bot.UpdMessage:
			return upd.Message != nil
		}
		return false
	}
}
