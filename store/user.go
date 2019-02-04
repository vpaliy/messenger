package store

import "github.com/vpaliy/telex/model"

type UserStore interface {
	Get(query Query) (*model.User, error)
	GetAll(query Query) ([]*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(user *model.User) error
}
