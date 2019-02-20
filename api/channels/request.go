package channels

import (
	"github.com/vpaliy/telex/api"
	"github.com/vpaliy/telex/model"
	"github.com/vpaliy/telex/utils"
)

type createChannelRequest struct {
	api.Binder
	Channel     string   `json:"channel" validate:"required"`
	Tags        []string `json:"tags"`
	Image       string   `json:"image"`
	Description string   `json:"description"`
	Type        string   `json:"type" validate:"required"`
	Private     bool     `json:"private"`
}

type updateChannelRequest struct {
	api.Binder
	Channel     string   `json:"channel" validate:"required"`
	Tags        []string `json:"tags"`
	Image       *string  `json:"image"`
	Description *string  `json:"description"`
	Private     *bool    `json:"private"`
}

type channelAction struct {
	api.Binder
	Channel string `json:"channel" validate:"required"`
}

type channelSearchRequest struct {
	api.Binder
	Query  string          `json:"query" validate:"required"`
	Tags   string          `json:"tags"`
	Latest utils.Timestamp `json:"latest"`
	Oldest utils.Timestamp `json:"oldest"`
	Limit  int16           `json:"limit"`
}

type markChannelRequest struct {
	api.Binder
	Channel   string          `json:"channel" validate:"required"`
	Timestamp utils.Timestamp `json:"ts" validate:"required"`
}

func (c *createChannelRequest) toChannel(id uint) *model.Channel {
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

func (r *updateChannelRequest) update(channel *model.Channel) {
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
