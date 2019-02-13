package messages

import (
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/utils"
	_ "log"
	"net/http"
)

func (h *Handler) GetMessages(c echo.Context) error {
	return nil
}

func (h *Handler) PostMessage(c echo.Context) error {
	request := new(createMessageRequest)
	message, err := request.bind(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
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
