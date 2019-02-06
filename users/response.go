package users

import (
	"github.com/vpaliy/telex/model"
)

type userResponse struct {
	Email    string  `json:"email"`
	Username string  `json:"username"`
	FullName string  `json:"fullName"`
	Bio      string  `json:"bio"`
	Image    *string `json:"image"`
}

func newUserResponse(user *model.User) *userResponse {
	response := new(userResponse)
	response.Email = user.Email
	response.Username = user.Username
	response.FullName = user.FullName
	response.Bio = user.Bio
	response.Image = user.Image
	return response
}
