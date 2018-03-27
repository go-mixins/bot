package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type MessageOption func(*tgbotapi.MessageConfig)
type EditMessageOption func(*tgbotapi.EditMessageTextConfig)
type EditMessageCaptionOption func(*tgbotapi.EditMessageCaptionConfig)
