package messages

import (
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/utils"
)

type Handler struct {
}

func (h *Handler) Register(group *echo.Group) {
	chat := group.Group("/chat", utils.JWTMiddleware())
	chat.GET("", h.GetMessages)
	chat.POST("", h.PostMessage)
	chat.PUT("/:id", h.EditMessage)
	chat.DELETE("/:id", h.DeleteMessage)
}
