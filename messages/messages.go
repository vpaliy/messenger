package messages

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/store"
	"github.com/vpaliy/telex/utils"
)

type Handler struct {
	messageStore store.MessageStore
}

// TODO: use dependency injection for this
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		messageStore: &MessageStore{db},
	}
}

func (h *Handler) Register(group *echo.Group) {
	chat := group.Group("/chat", utils.JWTMiddleware())
	chat.GET("", h.GetMessages)
	chat.POST("", h.PostMessage)
	chat.PUT("/:id", h.EditMessage)
	chat.DELETE("/:id", h.DeleteMessage)
}
