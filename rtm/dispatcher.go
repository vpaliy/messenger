package rtm

import (
	"sync"
)

type channelEvent struct {
	token   string
	client  *Client
	request *ChannelRequest
}

type postEvent struct {
	token        string
	client       *Client
	request      *CreateMessageRequest
	subscription *Subscription
}

type loadEvent struct {
	token   string
	client  *Client
	request *FetchMessagesRequest
}

type Dispatcher interface {
	Join(channelEvent) error
	Leave(channelEvent) error
	Typing(channelEvent) error
	Load(loadEvent) error
	Post(postEvent) error
}

type dispatcher struct {
	repository Repository
	hubs       sync.Map
	mutex      sync.Mutex
}

func NewDispatcher(repository Repository) Dispatcher {
	return &dispatcher{repository: repository}
}

func (d *dispatcher) Post(event postEvent) error {
	ctx := Context{event.token, event.request}
	message, err := d.repository.PostMessage(ctx)
	if err == nil {
		response := &ResponseMessage{Send, message}
		event.subscription.broadcast <- response
		return nil
	}
	return err
}

func (d *dispatcher) fetchHub(id string) (*Hub, error) {
	if value, ok := d.hubs.Load(id); ok {
		return value.(*Hub), nil
	}
	channel, err := d.repository.FetchChannel(id)
	if err != nil {
		return nil, err
	}
	hub := NewHub(channel.Name)
	d.hubs.Store(id, hub)
	return hub, nil
}

func (d *dispatcher) Join(event channelEvent) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	// create/get the hub
	request := event.request
	hub, err := d.fetchHub(request.Channel)
	if err != nil {
		return err
	}
	// join the channel, fetch the subscription
	ctx := Context{event.token, request}
	sub, err := d.repository.JoinChannel(ctx)
	if err != nil {
		return err
	}
	// register the client
	if !hub.HasClient(event.client) {
		hub.register <- event.client
	}
	// send back the subscription
	response := ResponseMessage{Join, sub}
	event.client.JSON(&response)
	return nil
}

func (d *dispatcher) Leave(event channelEvent) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	// create/get the hub
	request := event.request
	hub, err := d.fetchHub(request.Channel)
	if err != nil {
		return err
	}
	ctx := Context{event.token, request}
	sub, err := d.repository.LeaveChannel(ctx)
	if err != nil {
		return err
	}
	if hub.HasClient(event.client) {
		hub.unregister <- event.client
	}
	response := ResponseMessage{Join, sub}
	event.client.JSON(&response)
	return nil
}

func (d *dispatcher) Typing(event channelEvent) error {
	// TODO: implement
	return nil
}

func (d *dispatcher) Load(event loadEvent) error {
	ctx := Context{event.token, event.request}
	messages, err := d.repository.FetchMessages(ctx)
	if err != nil {
		return err
	}
	response := &ResponseMessage{Load, messages}
	event.client.JSON(response)
	return nil
}
