package bot

import "context"

type ContextFuncs interface {
	From(context.Context) *User
	Text(context.Context) string
}
