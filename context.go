package bot

import "context"

type ContextFuncs interface {
	UserName(context.Context) string
	Text(context.Context) string
}
