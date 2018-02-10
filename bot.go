package bot

import "github.com/andviro/middleware"

type Bot interface {
	Driver
	Run() error
	On(middleware.Predicate, middleware.Handler)
	Use(...middleware.Middleware)
	Handle(middleware.Handler)
}

//go:generate moq -out mock/bot.go -pkg mock . Bot
