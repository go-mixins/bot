package telegram

import (
	"context"

	"github.com/go-mixins/bot"
)

type MessageMeta struct {
	ID int
}

type ChatMeta struct {
	ID int64
}

type UserMeta struct {
	ID    int
	IsBot bool
}

func (drv *Driver) Me(ctx context.Context) *bot.User {
	return user(&drv.api.Self)
}

func (drv *Driver) From(ctx context.Context) *bot.User {
	return user(drv.from(ctx))
}

func (drv *Driver) Chat(ctx context.Context) *bot.Chat {
	return chat(drv.chat(ctx))
}

func (drv *Driver) Msg(ctx context.Context) *bot.Message {
	return message(drv.message(ctx))
}

func (drv *Driver) NewChatMembers(ctx context.Context) (res []*bot.User) {
	msg := drv.message(ctx)
	if msg.NewChatMembers == nil {
		return
	}
	res = make([]*bot.User, len(*msg.NewChatMembers))
	for i, u := range *msg.NewChatMembers {
		res[i] = user(&u)
	}
	return
}
