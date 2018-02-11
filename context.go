package bot

import "context"

type ContextFuncs interface {
	From(context.Context) *User
	Msg(context.Context) *Message
	Chat(context.Context) *Chat
	Debug(context.Context) string
}
