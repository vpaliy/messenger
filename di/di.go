//+build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	_ "github.com/vpaliy/telex/api"
	"github.com/vpaliy/telex/api/channels"
	"github.com/vpaliy/telex/api/messages"
	"github.com/vpaliy/telex/api/users"
	"github.com/vpaliy/telex/rtm"
	store "github.com/vpaliy/telex/store/gorm"
)

func NewRepository() *rtm.TestRepository {
	return &rtm.TestRepository{}
}

func InitializeChannelManager() rtm.ChannelManager {
	//  wire.Build(NewRepository)
	return rtm.NewChannelManager(NewRepository())
}

func InitializeChannelHandler(database *gorm.DB) *channels.Handler {
	wire.Build(
		channels.NewHandler,
		store.NewSubscriptionStore,
		store.NewChannelStore,
		store.NewUserStore,
	)
	return nil
}

func InitializeUserHandler(database *gorm.DB) *users.Handler {
	wire.Build(users.NewHandler, store.NewUserStore)
	return nil
}

func InitializeMessageHandler(database *gorm.DB) *messages.Handler {
	wire.Build(
		messages.NewHandler,
		store.NewMessageStore,
		store.NewChannelStore,
	)
	return nil
}
