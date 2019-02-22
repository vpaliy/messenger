package channels

import (
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/model"
	"github.com/vpaliy/telex/store"
	"github.com/vpaliy/telex/utils"
	_ "log"
	"net/http"
)

func (h *Handler) fetchChannel(c echo.Context) (*model.Channel, error) {
	request := new(ChannelRequest)
	if err := request.Bind(c); err != nil {
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

func (h *Handler) CreateChannel(c echo.Context) error {
	request := new(CreateChannelRequest)
	if err := request.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	channel := request.ToChannel(utils.GetUser(c).ID)
	if err := h.channelStore.Create(channel); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, NewChannelResponse(channel))
}

func (h *Handler) UpdateChannel(c echo.Context) error {
	channel, err := h.fetchChannel(c)
	// no channel, exit
	if channel == nil {
		return err
	}
	// TODO: allow adding admin users, so they can edit this channel as well
	if channel.IsCreator(utils.GetUser(c).ID) {
		return c.JSON(http.StatusForbidden, utils.Forbidden())
	}
	request := new(UpdateChannelRequest)
	if err := request.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	request.UpdateModel(channel)
	if err := h.channelStore.Update(channel); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, NewChannelResponse(channel))
}

func (h *Handler) FetchChannelInfo(c echo.Context) error {
	channel, err := h.fetchChannel(c)
	if channel == nil {
		return err
	}
	return c.JSON(http.StatusOK, NewChannelResponse(channel))
}

func (h *Handler) FetchChannels(c echo.Context) error {
	user, err := h.userStore.Fetch(utils.GetUser(c).Username)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	// fetch all channels that user has
	channels, err := h.channelStore.GetForMember(user.ID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, NewUserChannelsResponse(user, channels))
}

func (h *Handler) FetchSubscriptions(c echo.Context) error {
	user, err := h.userStore.Fetch(utils.GetUser(c).Username)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	subscriptions, err := h.subscriptionStore.FetchAll(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, NewUserSubscriptionsResponse(user, subscriptions))
}

func (h *Handler) JoinChannel(c echo.Context) error {
	channel, err := h.fetchChannel(c)
	if channel == nil {
		return err
	}
	user := utils.GetUser(c)
	subscription := channel.CreateSubscription(user.ID)
	if err := h.subscriptionStore.Create(channel, subscription); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, NewSubscriptionResponse(subscription))
}

func (h *Handler) LeaveChannel(c echo.Context) error {
	return nil
}

func (h *Handler) SearchChannels(c echo.Context) error {
	request := new(ChannelSearchRequest)
	if err := request.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	channels, err := h.channelStore.Search(
		request.Query,
		store.From(request.Oldest.Time()),
		store.To(request.Latest.Time()),
		store.Limit(request.Limit),
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, NewChannelsResponse(channels))
}

func (h *Handler) ArchiveChannel(c echo.Context) error {
	channel, err := h.fetchChannel(c)
	// no channel, exit
	if channel == nil {
		return err
	}
	if channel.IsCreator(utils.GetUser(c).ID) {
		return c.JSON(http.StatusForbidden, utils.Forbidden())
	}
	channel.Archived = true
	if err := h.channelStore.Update(channel); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.NoContent(http.StatusOK)
}

func (h *Handler) MarkSubscription(c echo.Context) error {
	// TODO: implement
	return nil
}

func (h *Handler) FetchSubscription(c echo.Context) error {
	// TODO: implement
	return nil
}
