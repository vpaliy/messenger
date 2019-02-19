package users

import (
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/store"
)

type Handler struct {
	userStore store.UserStore
}

func NewHandler(us store.UserStore) *Handler {
	return &Handler{
		userStore: us,
	}
}

func (h *Handler) Register(group *echo.Group) {
	group.POST("/users/login", h.Login)
	group.POST("/users/register", h.SignUp)
	group.POST("/users/forgot-password", h.ForgotPassword)
}
