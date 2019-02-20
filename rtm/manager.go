package rtm

import (
	"log"
	"sync"
)

type bundle struct {
	client  *Client
	channel string
}

type Dispatcher interface {
	Join(string, *Client)
	Leave(string, *Client)
	Post(CreateMessageRequest, *Client) (Content, error)
	Run()
}

type dispatcher struct {
	repository Repository
	channels   sync.Map
	mutex      sync.Mutex
	join       chan bundle
	leave      chan bundle
}

func (d *dispatcher) fetchHub(id string) (*Hub, error) {
	if value, ok := d.channels.Load(id); ok {
		return value.(*Channel), nil
	}
	value, err := d.repository.FetchChannel(id)
	if err != nil {
		return nil, err
	}
	d.channels.Store(id, value)
	return value, nil
}

func (d *dispatcher) Post(req *CreateMessageRequest, c *Client) (Content, error) {
	tr := TokenizedPostRequest{c.Token, req}
	return d.manager.PostMessage(&tr)
}

func (d *dispatcher) Join(req *ChannelRequest, c *Client) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	hub, err := d.fetchHub(req.Channel)
	if err != nil {
		client.JSONError(err)
	}
	if !hub.HasClient(c) {
		d.repository.Join(&TokenizedChannelRequest{c.Token, req})
	}
	//d.join <- bundle{client, channel}
}

func (d *dispatcher) Leave(channel string, client *Client) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	channel, err := d.fetchChannel(req.Channel)
	if err != nil {
		client.JSONError(err)
	}
	if !channel.HasClient(c) {
		d.repository.Join(&TokenizedChannelRequest{c.Token, req})
	}
	d.join <- bundle{client, channel}
}

func (d *dispatcher) Run() {
	for {
		select {
		case event := <-d.join:
			channel, err := d.fetchChannel(event.channel)
			if err != nil {
				event.client.JSONError(err)
				continue
			}
			channel.register <- event.client
		case event := <-d.leave:
			channel, err := d.fetchChannel(event.channel)
			if err != nil {
				event.client.JSONError(err)
				continue
			}
			channel.unregister <- event.client
		}
	}
}

func NewDispatcher(repository Repository) Dispatcher {
	return &dispatcher{
		repository: repository,
		join:       make(chan bundle),
		leave:      make(chan bundle),
	}
}
