package store

import "github.com/vpaliy/telex/model"

type MessageStore interface {
	Get(query Query) (*model.Message, error)
	GetAll(query Query) ([]*model.Message, error)
	Create(message *model.Message) error
	Update(message *model.Message) error
	Delete(message *model.Message) error
}
