package telegram

import (
	"bytes"
	"io/ioutil"
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

func newTransport() *transport {
	return &transport{
		chatLimits:   make(map[string]ratelimit.Limiter),
		globalLimit:  ratelimit.New(30),
		RoundTripper: http.DefaultTransport,
	}
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	req.Body = ioutil.NopCloser(bytes.NewReader(data))
	uid := req.FormValue("chat_id")
	req.Body = ioutil.NopCloser(bytes.NewReader(data))
	if uid != "" {
		var limiter ratelimit.Limiter
		var ok bool
		func() {
			t.l.Lock()
			defer t.l.Unlock()
			if limiter, ok = t.chatLimits[uid]; !ok {
				limiter = ratelimit.New(1, ratelimit.WithoutSlack)
				t.chatLimits[uid] = limiter
			}
		}()
		limiter.Take()
	}
	t.globalLimit.Take()
	return t.RoundTripper.RoundTrip(req)
}
