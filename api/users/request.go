package users

import (
	"github.com/vpaliy/telex/api"
	"github.com/vpaliy/telex/model"
)

type userLoginRequest struct {
	api.Binder
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type userRegisterRequest struct {
	api.Binder
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"fullName"`
	Bio      string `json:"bio"`
}

type forgotPasswordRequest struct {
	// either a username or email
	api.Binder
	Identifier string `json:"identifier" validate:"required"`
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
