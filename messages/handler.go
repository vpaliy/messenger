package messages

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/channels"
	"github.com/vpaliy/telex/store"
	"github.com/vpaliy/telex/utils"
)

type Handler struct {
	channelStore store.ChannelStore
	messageStore store.MessageStore
}

// TODO: use dependency injection for this
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		channelStore: channels.NewChannelStore(db),
		messageStore: &MessageStore{db},
	}
}

func (h *Handler) Register(group *echo.Group) {
	chat := group.Group("/chat")
	chat.Use(utils.JWTMiddleware())

	chat.GET(".list", h.GetMessages)
	chat.POST(".create", h.PostMessage)
	chat.POST(".edit", h.EditMessage)
	chat.POST(".delete", h.DeleteMessage)
}
