package store

import "github.com/vpaliy/telex/model"

type RoomStore interface {
	Get(query Query) (*model.Channel, error)
	GetAll(query Query) ([]*model.Channel, error)
	Create(room *model.Channel) error
	Update(room *model.Channel) error
	Delete(room *model.Channel) error
}
