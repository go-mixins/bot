package bot

import (
	"context"

	"github.com/andviro/middleware"
)

type Driver interface {
	Actions
	Predicates
	Context(context.Context, middleware.Handler) error
	Close() error
	Next() bool
}
