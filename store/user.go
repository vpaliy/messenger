package store

import "github.com/vpaliy/telex/model"

type UserStore interface {
	Get(Query) (*model.User, error)
	GetAll(Query) ([]*model.User, error)
	Create(*model.User) error
	Update(*model.User) error
	Delete(*model.User) error
}
