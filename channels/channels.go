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
	group.GET("/channels.info", h.FetchChannel)
	group.POST("/channels.create", h.CreateChannel)
	group.PUT("/channels.update", h.UpdateChannel)
	group.POST("/channels.archive", h.ArchiveChannel)
	group.DELETE("/channels.kick", h.KickUser)
	group.POST("/channels.invite", h.InviteUser)

	// set routes for channel subscriptions
	group.GET("/channels/:name/subscription", h.GetSubscription)
	group.POST("/channels/:name/subscription", h.JoinChannel)
}
