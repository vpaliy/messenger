package messages

import (
	"github.com/labstack/echo"
)

type Handler struct {
}

func (h *Handler) Register(group *echo.Group) {
	group.GET("/messages", h.GetMessages)
	group.POST("/messages", h.PostMessage)
	group.PUT("/messages", h.EditMessage)
	group.DELETE("/messages", h.DeleteMessage)
}
