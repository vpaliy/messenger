package users

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/store"
)

type Handler struct {
	userStore store.UserStore
}

func NewHandler(db *gorm.DB) *Handler {
	store := &UserStore{db}
	return &Handler{store}
}

func (h *Handler) Register(group *echo.Group) {
	group.POST("/users/login", h.Login)
	group.POST("/users/register", h.SignUp)
}
