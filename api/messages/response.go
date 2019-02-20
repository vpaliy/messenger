package messages

import (
	"github.com/vpaliy/telex/model"
	"github.com/vpaliy/telex/utils"
)

type Message struct {
	Text        string          `json:"text"`
	Username    string          `json:"username"`
	CreatedAt   utils.Timestamp `json:"timestamp"`
	Attachments []Attachment    `json:"attachments"`
}

type CreateMessageResponse struct {
	Channel string  `json:"channel"`
	Message Message `json:"message"`
}

type FetchMessagesResponse struct {
	Channel  string    `json:"channel"`
	Messages []Message `json:"messages"`
}

func NewFetchMessagesResponse(channel *model.Channel, messages []*model.Message) *FetchMessagesResponse {
	response := new(FetchMessagesResponse)
	response.Channel = string(channel.ID)
	response.Messages = make([]Message, len(messages))
	for i, m := range messages {
		response.Messages[i] = *NewMessageResponse(m)
	}
	return response
}

func NewMessageResponse(m *model.Message) *Message {
	response := new(Message)
	response.Text = m.Text
	response.Username = m.User.Username
	response.CreatedAt = utils.Timestamp(m.CreatedAt)
	response.Attachments = make([]Attachment, len(m.Attachments))
	for i, a := range m.Attachments {
		response.Attachments[i] = Attachment{
			Title:    a.Title,
			Text:     a.Text,
			ImageURL: a.ImageURL,
			AudioURL: a.AudioURL,
			VideoURL: a.VideoURL,
			ThumbURL: a.ThumbURL,
			Color:    a.Color,
		}
	}
	return response
}

func NewCreateMessageResponse(m *model.Message) *CreateMessageResponse {
	response := new(CreateMessageResponse)
	response.Message = *NewMessageResponse(m)
	response.Channel = string(m.ChannelID)
	return response
}
