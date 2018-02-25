package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Option func(*tgbotapi.MessageConfig)
