package bot

import (
	"context"

	"github.com/andviro/middleware"
)

type Driver interface {
	Actions
	Predicates
	ContextFuncs
	Context(context.Context, middleware.Handler) error
	Close() error
	Next() bool
}

//go:generate moq -out mock/driver.go -pkg mock . Driver
