package messages

import (
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/store"
	"github.com/vpaliy/telex/utils"
)

type Handler struct {
	channelStore store.ChannelStore
	messageStore store.MessageStore
}

func NewHandler(cs store.ChannelStore, ms store.MessageStore) *Handler {
	return &Handler{
		channelStore: cs,
		messageStore: ms,
	}
}

func (h *Handler) Register(group *echo.Group) {
	chat := group.Group("/chat")
	chat.Use(utils.JWTMiddleware())

	chat.GET(".list", h.GetMessages)
	chat.POST(".create", h.PostMessage)
	chat.POST(".edit", h.EditMessage)
	chat.POST(".search", h.Search)
	chat.POST(".delete", h.DeleteMessage)
}
