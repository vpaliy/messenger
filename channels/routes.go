package channels

import (
	"github.com/labstack/echo"
)

func (h *Handler) FetchChannel(c echo.Context) error {
	id := c.Param("id")
	query := store.NewQuery(map[string]interface{}{"id": id})
	channel, err := h.channelStore.Get(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if channel == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	return c.JSON(http.StatusOK, newChannelResponse(channel))
}

func (h *Handler) CreateChannel(c echo.Context) error {
	request := new(createChannelRequest)
	if err := request.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	channel := request.toChannel()
	if err := h.channelStore.Create(channel); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, newChannelResponse(channel))
}

func (h *Handler) UpdateChannel(c echo.Context) error {
	request := new(updateChannelRequest)
	if err := request.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	channel := request.toChannel()
	if err := h.channelStore.Update(channel); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, newChannelResponse(channel))
}

func (h *Handler) ArchiveChannel(c echo.Context) error {
	// TODO: implement
	return nil
}

func (h *Handler) GetSubscription(c echo.Context) error {
	// TODO: implement
	return nil
}

func (h *Handler) UpdateSubscription(c echo.Context) error {
	// TODO: implement
	return nil
}

func (h *Handler) JoinChannel(c echo.Context) error {
	// TODO: implement
	return nil
}

func (h *Handler) InviteUser(c echo.Context) error {
	// TODO: implement
	return nil
}

func (h *Handler) KickUser(c echo.Context) error {
	// TODO: implement
	return nil
}
