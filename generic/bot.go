package generic

import (
	"context"
	"sync"

	"github.com/andviro/middleware"

	"github.com/go-mixins/bot"
)

type Bot struct {
	bot.Driver
	mw middleware.Middleware
	h  middleware.Handler
	l  sync.RWMutex
}

func New(driver bot.Driver) (res *Bot, err error) {
	switch {
	case driver == nil:
		err = bot.Errors.New("driver not configured")
		return
	}
	res = &Bot{
		Driver: driver,
		mw:     driver.Context,
		h: func(context.Context) error {
			return nil
		},
	}
	return
}

func (b *Bot) Run() error {
	for b.Driver.Next() {
		if err := b.processUpdate(); err != nil {
			return err
		}
	}
	return nil
}

func (b *Bot) processUpdate() (err error) {
	b.l.RLock()
	defer b.l.RUnlock()
	return middleware.Middleware(b.Context).Use(b.mw)(context.Background(), b.h)
}

func (b *Bot) On(p middleware.Predicate, h middleware.Handler) {
	b.l.Lock()
	defer b.l.Unlock()
	b.mw = b.mw.On(p, h)
}

func (b *Bot) Use(mws ...middleware.Middleware) {
	b.l.Lock()
	defer b.l.Unlock()
	b.mw = b.mw.Use(mws...)
}
