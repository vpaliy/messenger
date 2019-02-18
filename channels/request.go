package channels

import (
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/model"
	"github.com/vpaliy/telex/utils"
)

type createChannelRequest struct {
	Channel     string   `json:"channel" validate:"required"`
	Tags        []string `json:"tags"`
	Image       string   `json:"image"`
	Description string   `json:"description"`
	Type        string   `json:"type" validate:"required"`
	Private     bool     `json:"private"`
}

type updateChannelRequest struct {
	Channel     string   `json:"channel" validate:"required"`
	Tags        []string `json:"tags"`
	Image       *string  `json:"image"`
	Description *string  `json:"description"`
	Private     *bool    `json:"private"`
}

type channelAction struct {
	Channel string `json:"channel" validate:"required"`
}

type channelSearchRequest struct {
	Query string `json:"query" validate:"required"`
	Tags  string `json:"tags"`
}

type markChannelRequest struct {
	Channel   string          `json:"channel" validate:"required"`
	Timestamp utils.Timestamp `json:"ts" validate:"required"`
}

// TODO: get rid of all these methods here, generalize it
func (r *createChannelRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

func (r *channelSearchRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

func (r *updateChannelRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

func (r *channelAction) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

func (r *markChannelRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
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
