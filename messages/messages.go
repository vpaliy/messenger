package messages

import (
	"github.com/labstack/echo"
)

type Handler struct {
}

func (h *Handler) Register(group *echo.Group) {
	group.GET("/chat", h.GetMessages)
	group.POST("/chat", h.PostMessage)
	group.PUT("/chat", h.EditMessage)
	group.DELETE("/chat", h.DeleteMessage)
}
