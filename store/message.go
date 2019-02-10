package store

import "github.com/vpaliy/telex/model"

type MessageStore interface {
	Get(Query) (*model.Message, error)
	GetAll(Query) ([]*model.Message, error)
	Create(*model.Message) error
	Update(*model.Message) error
	Delete(*model.Message) error
}
