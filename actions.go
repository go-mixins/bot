package bot

import "context"

type Actions interface {
	Reply(context.Context, interface{}) error
}
