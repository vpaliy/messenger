package channels

import (
	"github.com/vpaliy/telex/api"
	"github.com/vpaliy/telex/model"
	"github.com/vpaliy/telex/utils"
)

type CreateChannelRequest struct {
	api.Binder
	Channel     string   `json:"channel" validate:"required"`
	Tags        []string `json:"tags"`
	Image       string   `json:"image"`
	Description string   `json:"description"`
	Type        string   `json:"type" validate:"required"`
	Private     bool     `json:"private"`
}

type UpdateChannelRequest struct {
	api.Binder
	Channel     string   `json:"channel" validate:"required"`
	Tags        []string `json:"tags"`
	Image       *string  `json:"image"`
	Description *string  `json:"description"`
	Private     *bool    `json:"private"`
}

type ChannelRequest struct {
	api.Binder
	Channel string `json:"channel" validate:"required"`
}

type ChannelSearchRequest struct {
	api.Binder
	Query  string          `json:"query" validate:"required"`
	Tags   string          `json:"tags"`
	Latest utils.Timestamp `json:"latest"`
	Oldest utils.Timestamp `json:"oldest"`
	Limit  int16           `json:"limit"`
}

type MarkChannelRequest struct {
	api.Binder
	Channel   string          `json:"channel" validate:"required"`
	Timestamp utils.Timestamp `json:"ts" validate:"required"`
}

type Attachment struct {
	api.Binder
	Title    string `json:"title"`
	Text     string `json:"text"`
	ImageURL string `json:"image_url"`
	AudioURL string `json:"audio_url"`
	VideoURL string `json:"video_url"`
	ThumbURL string `json:"thumb_url"`
	Color    string `json:"color"`
}

type CreateMessageRequest struct {
	api.Binder
	Channel     string       `json:"channel" validate:"required"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

type FetchMessagesRequest struct {
	api.Binder
	Channel string          `json:"channel" validate:"required"`
	Latest  utils.Timestamp `json:"latest"`
	Oldest  utils.Timestamp `json:"oldest"`
	Limit   int16           `json:"limit"`
}

type EditMessageRequest struct {
	api.Binder
	ID      string `json:"message_id" validate:"required"`
	Channel string `json:"channel" validate:"required"`
	Text    string `json:"text" validate:"required"`
}

type SearchMessagesRequest struct {
	api.Binder
	Channel string          `json:"channel" validate:"required"`
	Query   string          `json:"query" validate:"required"`
	Latest  utils.Timestamp `json:"latest"`
	Oldest  utils.Timestamp `json:"oldest"`
	Limit   int16           `json:"limit"`
}

type DeleteMessageRequest struct {
	api.Binder
	ID string `json:"message_id" validate:"required"`
}

func (r *CreateMessageRequest) ToMessage(channel, user uint) *model.Message {
	message := new(model.Message)
	message.ChannelID = channel
	message.UserID = user
	message.Text = r.Text
	message.Attachments = make([]model.Attachment, len(r.Attachments))
	for i, a := range r.Attachments {
		message.Attachments[i] = model.Attachment{
			Title:    a.Title,
			Text:     a.Text,
			ImageURL: a.ImageURL,
			AudioURL: a.AudioURL,
			VideoURL: a.VideoURL,
			ThumbURL: a.ThumbURL,
			Color:    a.Color,
		}
	}
	return message
}

func (c *CreateChannelRequest) ToChannel(id uint) *model.Channel {
	return &model.Channel{
		CreatorID:   id,
		Name:        c.Channel,
		Tags:        model.CreateTags(c.Tags),
		Image:       c.Image,
		Description: c.Description,
		Type:        c.Type,
		Private:     c.Private,
	}
}

func (r *UpdateChannelRequest) UpdateModel(channel *model.Channel) {
	if len(r.Tags) > 0 {
		channel.Tags = model.CreateTags(r.Tags)
	}
	if r.Image != nil {
		channel.Image = *r.Image
	}
	if r.Description != nil {
		channel.Description = *r.Description
	}
	if r.Private != nil {
		channel.Private = *r.Private
	}
}
