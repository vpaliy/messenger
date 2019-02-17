package channels

import (
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/model"
	"github.com/vpaliy/telex/utils"
	_ "log"
	"net/http"
)

func (h *Handler) getChannel(c echo.Context) (*model.Channel, error) {
	request := new(channelAction)
	if err := request.bind(c); err != nil {
		return nil, c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	channel, err := h.channelStore.Fetch(request.Channel)
	if err != nil {
		return nil, c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if channel == nil {
		return nil, c.JSON(http.StatusNotFound, utils.NotFound())
	}
	return channel, nil
}

func (h *Handler) FetchChannel(c echo.Context) error {
	channel, err := h.getChannel(c)
	if channel == nil {
		return err
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
	channel, err := h.getChannel(c)
	if channel == nil {
		return err
	}
	if channel.IsCreator(utils.GetUserId(c)) {
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

func (h *Handler) FetchChannels(c echo.Context) error {
	user, err := h.userStore.Fetch(utils.GetUserId(c))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	channels, err := h.channelStore.GetForMember(user.ID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, newUserChannelsResponse(user, channels))
}

func (h *Handler) FetchSubscriptions(c echo.Context) error {
	user, err := h.userStore.Fetch(utils.GetUserId(c))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	subscriptions, err := h.subscriptionStore.FetchAll(user.ID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, newUserSubscriptionsResponse(user, subscriptions))
}

func (h *Handler) JoinChannel(c echo.Context) error {
	channel, err := h.getChannel(c)
	if channel == nil {
		return err
	}
	user, err := h.userStore.Fetch(utils.GetUserId(c))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	subscription := channel.CreateSubscription(user.ID)
	if err := h.subscriptionStore.Create(channel, subscription); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, newSubscriptionResponse(subscription))
}

func (h *Handler) MarkSubscription(c echo.Context) error {
	// TODO: implement
	return nil
}

func (h *Handler) ArchiveChannel(c echo.Context) error {
	// TODO: implement
	return nil
}

func (h *Handler) FetchSubscription(c echo.Context) error {
	// TODO: implement
	return nil
}

func (h *Handler) SearchChannels(e echo.Context) error {
	return nil
}
