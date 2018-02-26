package telegram

import (
	"fmt"
	"net/http"
	"sync"

	"go.uber.org/ratelimit"
)

type transport struct {
	l           sync.RWMutex
	chatLimits  map[string]ratelimit.Limiter
	globalLimit ratelimit.Limiter
	http.RoundTripper
}

func newTransport() http.RoundTripper {
	return &transport{
		chatLimits:   make(map[string]ratelimit.Limiter),
		globalLimit:  ratelimit.New(30),
		RoundTripper: http.DefaultTransport,
	}
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Printf("request: %+v\n", req)
	uid := req.PostFormValue("chat_id")
	if uid != "" {
		var limiter ratelimit.Limiter
		var ok bool
		func() {
			t.l.Lock()
			defer t.l.Unlock()
			if limiter, ok = t.chatLimits[uid]; !ok {
				limiter = ratelimit.New(1)
				t.chatLimits[uid] = limiter
			}
		}()
		limiter.Take()
	}
	t.globalLimit.Take()
	resp, err := t.RoundTripper.RoundTrip(req)
	fmt.Printf("%+v %+v", resp, err)
	return resp, err
}
