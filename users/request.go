package users

import (
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/model"
)

type userLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func newUserLoginRequest() *userLoginRequest {
	return new(userLoginRequest)
}

func (r *userLoginRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

type userRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"fullName"`
	Bio      string `json:"bio"`
}

func newUserRegisterRequest() *userRegisterRequest {
	return new(userRegisterRequest)
}

func (r *userRegisterRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

func (r *userRegisterRequest) toUser() *model.User {
	user := new(model.User)
	user.Username = r.Username
	user.Email = r.Email
	user.FullName = r.FullName
	user.Bio = r.Bio
	if err := user.SetPassword(r.Password); err != nil {
		return nil
	}
	return user
}
