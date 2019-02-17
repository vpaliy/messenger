package channels

import (
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/store"
	"github.com/vpaliy/telex/utils"
)

type Handler struct {
	userStore         store.UserStore
	channelStore      store.ChannelStore
	subscriptionStore store.SubscriptionStore
}

func NewHandler(cs store.ChannelStore, ss store.SubscriptionStore, us store.UserStore) *Handler {
	return &Handler{
		userStore:         us,
		channelStore:      cs,
		subscriptionStore: ss,
	}
}

func (h *Handler) Register(group *echo.Group) {
	channels := group.Group("/channels")
	subscriptions := group.Group("/subscriptions")

	channels.Use(utils.JWTMiddleware())
	subscriptions.Use(utils.JWTMiddleware())

	channels.POST(".create", h.CreateChannel)
	channels.POST(".update", h.UpdateChannel)
	channels.POST(".join", h.JoinChannel)
	channels.GET(".list", h.FetchChannels)
	channels.GET(".info", h.FetchChannel)
	channels.GET(".search", h.SearchChannels)

	subscriptions.GET(".list", h.FetchSubscriptions)
	subscriptions.GET(".info", h.FetchSubscription)
	subscriptions.POST(".mark", h.MarkSubscription)
	subscriptions.POST(".join", h.JoinChannel)
}
