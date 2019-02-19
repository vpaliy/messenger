package store

import "github.com/vpaliy/telex/model"

type Arg interface{}

type UserStore interface {
	Fetch(Arg) (*model.User, error)
	// search by name
	Search(Arg, ...Option) ([]*model.User, error)
	Create(*model.User) error
	Update(*model.User) error
	Delete(*model.User) error
}

type MessageStore interface {
	Fetch(Arg) (*model.Message, error)
	GetForChannel(Arg, ...Option) ([]*model.Message, error)
	GetForUser(Arg, ...Option) ([]*model.Message, error)
	Search(channel, user Arg, opts ...Option) ([]*model.Message, error)
	Create(*model.Message) error
	Update(*model.Message) error
	Delete(*model.Message) error
}

type SubscriptionStore interface {
	//fetch by channel's name
	Fetch(Arg) (*model.Subscription, error)
	//fetch all subscriptions for a user
	FetchAll(Arg, ...Option) ([]*model.Subscription, error)
	Create(*model.Channel, *model.Subscription) error
	Update(*model.Subscription) error
	Delete(*model.Channel, *model.Subscription) error
}

type ChannelStore interface {
	Fetch(Arg) (*model.Channel, error)
	Search(Arg, ...Option) ([]*model.Channel, error)
	GetForMember(Arg, ...Option) ([]*model.Channel, error)
	GetCreatedBy(Arg, ...Option) ([]*model.Channel, error)
	Create(*model.Channel) error
	Update(*model.Channel) error
	Delete(*model.Channel) error
}
