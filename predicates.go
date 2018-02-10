package bot

import (
	"github.com/andviro/middleware"
)

type Predicates interface {
	Command(string) middleware.Predicate
	Hears(string) middleware.Predicate
	Message(MessageType) middleware.Predicate
	Update(UpdateType) middleware.Predicate
}
