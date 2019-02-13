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
	channels.Use(utils.JWTMiddleware())
	channels.GET("/:id", h.FetchChannel)
	channels.POST("", h.CreateChannel)
	channels.PUT("/:id", h.UpdateChannel)
	channels.DELETE("/:id", h.ArchiveChannel)

	channels.GET("/subscriptions", h.GetAllSubscriptions)
	channels.GET("/:channel/subscriptions/:id", h.GetSubscription)
	channels.DELETE("/:channel/subscriptions/:id", h.KickUser)
	channels.POST("/:id/subscriptions", h.JoinChannel)
}
