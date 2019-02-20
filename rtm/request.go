package rtm

import (
	"github.com/mitchellh/mapstructure"
	"github.com/vpaliy/telex/api/channels"
	"github.com/vpaliy/telex/api/messages"
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

type (
	CreateMessageRequest messages.CreateMessageRequest
	FetchMessagesRequest messages.FetchMessagesRequest
	EditMessageRequest   messages.EditMessageRequest
	ChannelRequest       channels.ChannelRequest
)

type WSEvent struct {
	Event   string      `json:"event"`
	Request interface{} `json:"request"`
}

type HelloRequest struct {
	UserAgent string `json:"ua"`
	Version   string `json:"version"`
	DeviceId  string `json:"dev"`
	Platform  string `json:"platform"`
}

type TypingRequest struct {
	Channel string `json:"channel"`
}

func (event *WSEvent) Unmarshal(raw []byte) error {
	return json.Unmarshal(raw, event)
}

func (event *WSEvent) DecodeAction(action interface{}) error {
	return mapstructure.Decode(event.Action, action)
}
