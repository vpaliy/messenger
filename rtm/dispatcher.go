package rtm

import (
	"sync"
)

type channelEvent struct {
	client  *Client
	channel string
}

type postEvent struct {
	token   string
	request *CreateMessageRequest
}

type Dispatcher interface {
	Join(channelEvent) (Content, error)
	Leave(channelEvent) (Content, error)
	Post(postEvent) (Content, error)
}

type dispatcher struct {
	repository Repository
	hubs       sync.Map
	mutex      sync.Mutex
}

func NewDispatcher(repository Repository) Dispatcher {
	return &dispatcher{repository: repository}
}

func (d *dispatcher) fetchHub(id string) (*Hub, error) {
	if value, ok := d.hubs.Load(id); ok {
		return value.(*Hub), nil
	}
	value, err := d.repository.FetchChannel(id)
	if err != nil {
		return nil, err
	}
	d.hubs.Store(id, value)
	// TODO: create a new channel here
	return nil, nil
}

func (d *dispatcher) Post(event postEvent) (Content, error) {
	tr := TokenizedPostRequest{event.token, event.request}
	return d.repository.PostMessage(&tr)
}

func (d *dispatcher) Join(event channelEvent) (Content, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	hub, err := d.fetchHub(event.channel)
	if err != nil {
		return nil, err
	}
	if !hub.HasClient(event.client) {
		// TODO: join the user
	}
	hub.register <- event.client
}

func (d *dispatcher) Leave(event channelEvent) (Content, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	hub, err := d.fetchHub(event.channel)
	if err != nil {
		return nil, err
	}
	if !hub.HasClient(event.client) {
		// TODO:
	}
	hub.unregister <- event.client
}
