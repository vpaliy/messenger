package channels

import (
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/model"
	"github.com/vpaliy/telex/store"
	"github.com/vpaliy/telex/utils"
	_ "log"
	"net/http"
)

func (h *Handler) getChannel(c echo.Context, param string) (*model.Channel, error) {
	id := c.Param(param)
	query := store.NewQuery(map[string]interface{}{"id": id})
	return h.channelStore.Get(query)
}

func (h *Handler) FetchChannel(c echo.Context) error {
	channel, err := h.getChannel(c, "id")
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
	channel, err := request.bind(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := h.channelStore.Create(channel); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, newChannelResponse(channel))
}

func (h *Handler) UpdateChannel(c echo.Context) error {
	channel, err := h.getChannel(c, "id")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if channel == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	claims := utils.GetJWTClaims(c)
	if channel.CreatorID != claims.ID {
		return c.JSON(http.StatusForbidden, utils.Forbidden())
	}

	request := new(updateChannelRequest)
	if err := request.bind(c, channel); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := h.channelStore.Update(channel); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, newChannelResponse(channel))
}

func (h *Handler) ArchiveChannel(c echo.Context) error {
	// TODO: implement
	return nil
}

func (h *Handler) GetSubscription(c echo.Context) error {
	// TODO: implement
	return nil
}

func (h *Handler) GetAllSubscriptions(c echo.Context) error {
	claims := utils.GetJWTClaims(c)
	query := store.NewQuery(map[string]interface{}{"user_id": claims.ID})
	subscriptions, err := h.subscriptionStore.GetAll(query)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	user := &subscriptions[0].User
	return c.JSON(http.StatusOK, newUserSubscriptionsResponse(user, subscriptions))
}

func (h *Handler) UpdateSubscription(c echo.Context) error {
	// TODO: implement
	return nil
}

func (h *Handler) JoinChannel(c echo.Context) error {
	channel, err := h.getChannel(c, "id")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if channel == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	claims := utils.GetJWTClaims(c)
	subscription := channel.CreateSubscription(claims.ID)
	if err := h.subscriptionStore.Create(channel, subscription); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, newSubscriptionResponse(subscription))
}

func (h *Handler) InviteUser(c echo.Context) error {
	// TODO: implement
	return nil
}

func (h *Handler) KickUser(c echo.Context) error {
	// TODO: implement
	return nil
}
