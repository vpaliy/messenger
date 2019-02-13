package messages

import (
	"github.com/vpaliy/telex/model"
	"github.com/vpaliy/telex/utils"
	_ "log"
)

type message struct {
	Text        string          `json:"text"`
	Username    string          `json:"username"`
	CreatedAt   utils.Timestamp `json:"timestamp"`
	Attachments []attachment    `json:"attachments"`
}

type createMessageResponse struct {
	Channel string  `json:"channel"`
	Message message `json:"message"`
}

type fetchMessagesResponse struct {
	Channel  string    `json:"channel"`
	Messages []message `json:"messages"`
}

func newFetchMessagesResponse(channel string, messages []*model.Message) *fetchMessagesResponse {
	response := new(fetchMessagesResponse)
	response.Channel = channel
	response.Messages = make([]message, len(messages))
	for i, m := range messages {
		response.Messages[i] = *newMessageResponse(m)
	}
	return response
}

func newMessageResponse(m *model.Message) *message {
	response := new(message)
	response.Text = m.Text
	response.Username = m.User.Username
	response.CreatedAt = utils.Timestamp(m.CreatedAt)
	response.Attachments = make([]attachment, len(m.Attachments))
	for i, a := range m.Attachments {
		response.Attachments[i] = attachment{
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

func newCreateMessageResponse(m *model.Message) *createMessageResponse {
	response := new(createMessageResponse)
	response.Message = *newMessageResponse(m)
	response.Channel = string(m.ChannelID)
	return response
}
