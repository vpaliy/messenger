package rtm

import (
	"github.com/vpaliy/telex/model"
	"github.com/vpaliy/telex/store"
	"github.com/vpaliy/telex/utils"
)

type Content interface{}

type Context struct {
	Token   string
	Request interface{}
}

type Repository interface {
	FetchChannel(string) (*model.Channel, error)
	JoinChannel(Context) (Content, error)
	LeaveChannel(Context) (Content, error)
	PostMessage(Context) (Content, error)
}

type repository struct {
	userStore         store.UserStore
	messageStore      store.MessageStore
	subscriptionStore store.MessageStore
	channelStore      store.ChannelStore
}

func (r *repository) FetchChannel(ch string) (*model.Channel, error) {
	return r.channelStore.Fetch(ch)
}

func (r *repository) PostMessage(ctx Context) (Content, error) {
	request := ctx.Request.(*ChannelRequest)
	channel, err := r.channelStore.Fetch(request.Channel)
	// if there is an error
	if channel == nil || err != nil {
		return nil, err
	}
	// check if the user is subscribed to the chat
	currentUser := utils.GetUserFromToken(ctx.Token)
	if !channel.HasUser(currentUser.ID) {
		// TODO: return an error that the user is not here
		return nil, nil
	}
	// submit a message
	message := request.ToMessage(
		currentUser.ID, channel.ID,
	)
	// create message and send an error message if fails
	if err := r.messageStore.Create(message); err != nil {
		// TODO: notify that the message has failed to be written to the db
		return nil, err
	}
	return message, nil
}

func (r *repository) JoinChannel(ctx Context) (Content, error) {
	request := ctx.Request.(*ChannelRequest)
	channel, err := r.channelStore.Fetch(request.Channel)
	// notify about the error
	if channel == nil || err != nil {
		return nil, err
	}
	// get the user from JWT token
	user := utils.GetUserFromToken(ctx.Token)
	// create a subscription
	subscription := channel.CreateSubscription(user.ID)
	// create it
	if err := r.subscriptionStore.Create(channel, subscription); err != nil {
		return nil, err
	}
	return subscription, nil
}

func (r *repository) LeaveChannel(ctx Context) (Content, error) {
	request := ctx.Request.(*ChannelRequest)
	channel, err := r.channelStore.Fetch(request.Channel)
	// notify about the error
	if channel == nil || err != nil {
		return nil, err
	}
	// get the user from JWT token
	user := utils.GetUserFromToken(ctx.Token)
	// create a subscription
	subscription := channel.CreateSubscription(user.ID)
	// create it
	if err := r.subscriptionStore.Create(channel, subscription); err != nil {
		return nil, err
	}
	return subscription, nil
}
