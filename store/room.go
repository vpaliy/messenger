package store

import "github.com/vpaliy/telex/model"

type SubscriptionStore interface {
	Get(Query) (*model.Subscription, error)
	GetAll(Query) ([]*model.Subscription, error)
	Create(*model.Channel, *model.Subscription) error
	Update(*model.Channel, *model.Subscription) error
	Delete(*model.Channel, *model.Subscription) error
}

type ChannelStore interface {
	Get(Query) (*model.Channel, error)
	GetAll(Query) ([]*model.Channel, error)
	Create(*model.Channel) error
	Update(*model.Channel) error
	Delete(*model.Channel) error
}
