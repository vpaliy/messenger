package messages

import (
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/store"
	"github.com/vpaliy/telex/utils"
	_ "log"
	"net/http"
	"time"
)

func (h *Handler) GetMessages(c echo.Context) error {
	request := new(fetchMessagesRequest)
	if err := request.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	channel, err := h.channelStore.Fetch(request.Channel)
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
	messages, err := h.messageStore.GetAll(
		query,
		store.From(request.Oldest.Time()),
		store.To(request.Latest.Time()),
		store.Limit(request.Limit),
	)
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
	channel, err := h.channelStore.Fetch(request.Channel)
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
