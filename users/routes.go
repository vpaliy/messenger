package users

import (
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/store"
	"github.com/vpaliy/telex/utils"
	"net/http"
)

func (h *Handler) Login(c echo.Context) error {
	request := newUserLoginRequest()
	if err := request.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	query := store.NewQuery(map[string]interface{}{
		"email":    request.Username,
		"username": request.Username,
	})
	user, err := h.userStore.Get(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	if !user.CheckPasswordHash(request.Password) {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	return c.JSON(http.StatusOK, newUserResponse(user))
}

func (h *Handler) SignUp(c echo.Context) error {
	request := newUserRegisterRequest()
	if err := request.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	user := request.toUser()
	if err := h.userStore.Create(user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, newUserResponse(user))
}

func (h *Handler) ForgotPassword(c echo.Context) error {
	return nil
}
