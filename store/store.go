package store

import "github.com/vpaliy/telex/model"

type UserStore interface {
	Fetch(string) (*model.User, error)
	// search by name
	Search(string, ...Option) ([]*model.User, error)
	Create(*model.User) error
	Update(*model.User) error
	Delete(*model.User) error
}

type MessageStore interface {
	GetForChannel(string, ...Option) ([]*model.Message, error)
	GetForUser(string, ...Option) ([]*model.Message, error)
	Create(*model.Message) error
	Update(*model.Message) error
	Delete(*model.Message) error
}

type SubscriptionStore interface {
	//fetch by channel's name
	Fetch(string) (*model.Subscription, error)
	//fetch all subscriptions for a user
	FetchAll(string, ...Option) ([]*model.Subscription, error)
	Create(*model.User, *model.Channel) error
	Update(*model.Subscription) error
	Delete(*model.Subscription) error
}

type ChannelStore interface {
	Fetch(string) (*model.Channel, error)
	Search(string, ...Option) ([]*model.Channel, error)
	GetForMember(string, ...Option) ([]*model.Channel, error)
	GetCreatedBy(string, ...Option) ([]*model.Channel, error)
	Create(*model.Channel) error
	Update(*model.Channel) error
	Delete(*model.Channel) error
}
