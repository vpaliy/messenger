package messages

import (
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/model"
	"github.com/vpaliy/telex/store"
	"github.com/vpaliy/telex/utils"
	_ "log"
	"net/http"
	"time"
)

func (h *Handler) getChannel(channel string) (*model.Channel, error) {
	query := store.NewQuery().Append("id", channel).
		AddPreload("Creator").
		AddPreload("Members")
	return h.channelStore.Get(query)
}

func (h *Handler) GetMessages(c echo.Context) error {
	request := new(fetchMessagesRequest)
	if err := request.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	channel, err := h.getChannel(request.Channel)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if channel == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	claims := utils.GetJWTClaims(c)
	if !channel.HasUser(claims.ID) {
		return c.JSON(http.StatusForbidden, utils.Forbidden())
	}
	query := store.NewQuery().
		Append("channel_id", channel.ID).
		AddPreload("Attachments").
		AddPreload("User")
	if request.Limit > 0 {
		query = query.SetLimit(request.Limit)
	}
	if !request.Oldest.IsZero() {
		if !request.Latest.IsZero() {
			query = query.SetTimeRange(
				request.Oldest.Time(),
				request.Latest.Time(),
			)
		} else {
			query = query.SetTimeRange(
				request.Oldest.Time(),
				time.Now(),
			)
		}
	}
	messages, err := h.messageStore.GetAll(query)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, newFetchMessagesResponse(string(channel.ID), messages))
}

func (h *Handler) PostMessage(c echo.Context) error {
	request := new(createMessageRequest)
	message, err := request.bind(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	channel, err := h.getChannel(request.Channel)
	// internal server error
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	// channel does not exist
	if channel == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	// check if the user is subscribed to the chat
	claims := utils.GetJWTClaims(c)
	if !channel.HasUser(claims.ID) {
		return c.JSON(http.StatusForbidden, utils.Forbidden())
	}
	// create message and send an error message if fails
	if err := h.messageStore.Create(message); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, newCreateMessageResponse(message))
}

func (h *Handler) DeleteMessage(c echo.Context) error {
	return nil
}

func (h *Handler) EditMessage(c echo.Context) error {
	return nil
}
