package telegram

import (
	"encoding/json"
)

// CallbackData is sent on inline button press
type CallbackData struct {
	Type string
	Data json.RawMessage
}

func (cbd CallbackData) String() string {
	s, _ := json.Marshal(&cbd)
	return string(s)
}
