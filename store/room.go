package store

import "github.com/vpaliy/telex/model"

type RoomStore interface {
	Get(query Query) (*model.Room, error)
	GetAll(query Query) ([]*model.Room, error)
	Create(room *model.Room) error
	Update(room *model.Room) error
	Delete(room *model.Room) error
}
