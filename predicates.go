package bot

import (
	"context"

	"github.com/andviro/middleware"
)

type Predicates interface {
	Command(string) middleware.Predicate
	Hears(string) middleware.Predicate
	Message(context.Context) bool
}
