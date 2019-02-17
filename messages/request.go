package messages

import (
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/model"
	"github.com/vpaliy/telex/utils"
	"strconv"
)

type attachment struct {
	Title    string `json:"title"`
	Text     string `json:"text"`
	ImageURL string `json:"image_url"`
	AudioURL string `json:"audio_url"`
	VideoURL string `json:"video_url"`
	ThumbURL string `json:"thumb_url"`
	Color    string `json:"color"`
}

type createMessageRequest struct {
	Channel     string       `json:"channel" validate:"required"`
	Text        string       `json:"text"`
	Attachments []attachment `json:"attachments"`
}

type fetchMessagesRequest struct {
	Channel string          `json:"channel" validate:"required"`
	Latest  utils.Timestamp `json:"latest"`
	Oldest  utils.Timestamp `json:"oldest"`
	Limit   int             `json:"limit"`
}

type editMessageRequest struct {
	ID      string `json:"message_id" validate:"required"`
	Channel string `json:"channel" validate:"required"`
	Text    string `json:"text" validate:"required"`
}

type searchMessagesRequest struct {
	Channel string          `json:"channel" validate:"required"`
	Query   string          `json:"query" validate:"required"`
	Latest  utils.Timestamp `json:"latest"`
	Oldest  utils.Timestamp `json:"oldest"`
	Limit   int             `json:"limit"`
}

type deleteMessageRequest struct {
	ID string `json:"message_id" validate:"required"`
}

func (r *deleteMessageRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

func (r *fetchMessagesRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

func (r *createMessageRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

func (r *editMessageRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

func (r *searchMessagesRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

func (r *createMessageRequest) createMessage(channel, user uint) *model.Message {
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
