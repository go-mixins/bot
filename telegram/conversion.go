package telegram

import (
	"github.com/go-mixins/bot"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func message(msg *tgbotapi.Message) *bot.Message {
	if msg == nil {
		return nil
	}
	return &bot.Message{
		Text: msg.Text,
		Meta: bot.Meta{
			Platform: "telegram",
			Data: MessageMeta{
				ID: msg.MessageID,
			},
		},
	}
}

func user(u *tgbotapi.User) *bot.User {
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
			Data: UserMeta{
				ID:    u.ID,
				IsBot: u.IsBot,
			},
		},
	}
}

func chat(chat *tgbotapi.Chat) *bot.Chat {
	if chat == nil {
		return nil
	}
	return &bot.Chat{
		Type:        chat.Type,
		Title:       chat.Title,
		UserName:    chat.UserName,
		FirstName:   chat.FirstName,
		LastName:    chat.LastName,
		Description: chat.Description,
		Meta: bot.Meta{
			Platform: "telegram",
			Data: ChatMeta{
				ID: chat.ID,
			},
		},
	}
}
