package messages

import (
	"github.com/labstack/echo"
)

type Handler struct {
}

func (h *Handler) Register(group *echo.Group) {
	group.GET("/chat", h.GetMessages)
	group.POST("/chat", h.PostMessage)
	group.PUT("/chat/:id", h.EditMessage)
	group.DELETE("/chat/:id", h.DeleteMessage)
}
