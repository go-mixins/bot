package bot

import (
	"context"
)

type ContextFuncs interface {
	Me(context.Context) *User
	From(context.Context) *User
	Msg(context.Context) *Message
	Chat(context.Context) *Chat
	NewChatMembers(context.Context) []*User
}
