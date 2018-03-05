package bot

import (
	"github.com/andviro/middleware"
)

type Bot interface {
	On(middleware.Predicate, middleware.Handler, ...middleware.Middleware)
	Use(...middleware.Middleware)
}

//go:generate moq -out mock/bot.go -pkg mock . Bot
