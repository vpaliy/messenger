package rtm

import (
	"github.com/mitchellh/mapstructure"
)

const (
	Join    = "join"
	Leave   = "leave"
	Send    = "send"
	Hello   = "hello"
	Typing  = "typing"
	Load    = "load"
	Goodbye = "goodbye"
)

type ActionRequest struct {
	Event  string      `json:"event"`
	Action interface{} `json:"action"`
}

type LoadRequest struct {
	From    string `json:"id"`
	Channel string `json:"channel"`
	Limit   int    `json:"limit,omitempty"`
}

type JoinAction struct {
	From    string                 `json:"id"`
	Channel string                 `json:"channel"`
	Meta    map[string]interface{} `json:"meta"`
}

type MessageAction struct {
	From    string                 `json:"id"`
	Channel string                 `json:"channel"`
	Meta    map[string]interface{} `json:"meta"`
	Content interface{}            `json:"content"`
}

type HelloAction struct {
	From      string `json:"id"`
	UserAgent string `json:"ua"`
	Version   string `json:"version"`
	DeviceId  string `json:"dev"`
	Platform  string `json:"platform"`
}

type TypingAction struct {
	From    string `json:"id"`
	Channel string `json:"channel"`
}

func (a *ActionRequest) DecodeAction() (interface{}, error) {
	var value interface{}
	err := mapstructure.Decode(a.Action, value)
	return value, err
}

// TODO: find a better way to do this
func (a *JoinAction) Decode(action *ActionRequest) (*JoinAction, error) {
	err := mapstructure.Decode(action.Action, a)
	return a, err
}

func (a *MessageAction) Decode(action *ActionRequest) (*MessageAction, error) {
	err := mapstructure.Decode(action.Action, a)
	return a, err
}

func (a *HelloAction) Decode(action *ActionRequest) (*HelloAction, error) {
	err := mapstructure.Decode(action.Action, a)
	return a, err
}

func (a *TypingAction) Decode(action *ActionRequest) (*TypingAction, error) {
	err := mapstructure.Decode(action.Action, a)
	return a, err
}

func (a *LoadRequest) Decode(action *ActionRequest) (*LoadRequest, error) {
	err := mapstructure.Decode(action.Action, a)
	return a, err
}
