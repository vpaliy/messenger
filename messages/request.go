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

func (r *fetchMessagesRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

func (r *createMessageRequest) bind(c echo.Context) (*model.Message, error) {
	if err := c.Bind(r); err != nil {
		return nil, err
	}
	if err := c.Validate(r); err != nil {
		return nil, err
	}
	claims := utils.GetJWTClaims(c)
	message := r.toMessage()
	message.UserID = claims.ID
	return message, nil
}

func (r *createMessageRequest) toMessage() *model.Message {
	id, _ := strconv.Atoi(r.Channel)
	message := new(model.Message)
	message.ChannelID = uint(id)
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
