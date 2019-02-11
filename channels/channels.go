package channels

import (
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/store"
)

type Handler struct {
	channelStore      ChannelStore
	subscriptionStore SubscriptionStore
}

// TODO: use dependency injection for this
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		channelStore:      &ChannelStore{db},
		subscriptionStore: &SubscriptionStore{db},
	}
}

func (h *Handler) Register(group *echo.Group) {
	group.GET("/channels/:id", h.FetchChannel)
	group.POST("/channels", h.CreateChannel)
	group.PUT("/channels/:id", h.UpdateChannel)
	group.DELETE("/channels/:id", h.ArchiveChannel)

	group.GET("/channels/:id/subscriptions/:id", h.GetSubscription)
	group.DELETE("channels/:id/subscriptions/:id", h.KickUser)
	group.POST("/channels/:id/subscriptions", h.JoinChannel)
}
