package channels

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/store"
	"github.com/vpaliy/telex/utils"
)

type Handler struct {
	channelStore      store.ChannelStore
	subscriptionStore store.SubscriptionStore
}

// TODO: use dependency injection for this
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		channelStore:      &ChannelStore{db},
		subscriptionStore: &SubscriptionStore{db},
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
	channels.POST(".kick", h.KickUser)
	channels.GET(".list", h.FetchChannels)
	channels.GET(".info", h.FetchChannel)

	subscriptions.GET(".list", h.FetchSubscriptions)
	subscriptions.GET(".info", h.FetchSubscription)
	subscriptions.POST(".mark", h.MarkSubscription)
	subscriptions.POST(".join", h.JoinChannel)
}
