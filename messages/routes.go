package messages

import (
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/model"
	"github.com/vpaliy/telex/store"
	"github.com/vpaliy/telex/utils"
	_ "log"
	"net/http"
)

func (h *Handler) getChannel(channel string) (*model.Channel, error) {
	query := store.NewQuery(map[string]interface{}{"id": channel})
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
	if channel.HasUser(claims.ID) {
		return c.JSON(http.StatusForbidden, utils.Forbidden())
	}
	query := store.NewQuery(map[string]interface{}{
		"channel_id": channel.ID,
	})
	messages, err := h.messageStore.GetAll(query)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, newFetchMessagesResponse(string(channel.ID), messages))
}

func (h *Handler) PostMessage(c echo.Context) error {
	request := new(createMessageRequest)
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
	if channel.HasUser(claims.ID) {
		return c.JSON(http.StatusForbidden, utils.Forbidden())
	}

	message := request.toMessage(c)
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
