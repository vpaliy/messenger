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

type markChannelRequest struct {
	Channel   string          `json:"channel" validate:"required"`
	Timestamp utils.Timestamp `json:"ts" validate:"required"`
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

func (r *updateChannelRequest) bind(c echo.Context, channel *model.Channel) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
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
	return nil
}

func (r *createChannelRequest) bind(c echo.Context) (*model.Channel, error) {
	if err := c.Bind(r); err != nil {
		return nil, err
	}
	if err := c.Validate(r); err != nil {
		return nil, err
	}
	//	claims := utils.GetJWTClaims(c)
	channel := r.toChannel()
	//channel.CreatorID = claims.ID
	return channel, nil
}

func (c *createChannelRequest) toChannel() *model.Channel {
	return &model.Channel{
		Name:        c.Channel,
		Tags:        model.CreateTags(c.Tags),
		Image:       c.Image,
		Description: c.Description,
		Type:        c.Type,
		Private:     c.Private,
	}
}
