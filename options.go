package bot

import "context"

type Option func(context.Context, ...interface{})

type Options interface {
	WithReply(context.Context, ...interface{})
}
