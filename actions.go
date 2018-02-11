package bot

import "context"

type Actions interface {
	Reply(context.Context, string, ...Option) error
}
