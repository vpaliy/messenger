package rtm

import (
	"log"
	"sync"
)

type clientEvent struct {
	client  *Client
	channel string
}

type ChannelManager interface {
	Join(channel string, client *Client)
	Leave(channel string, client *Client)
	Run()
}

type channelManager struct {
	repository Repository
	channels   sync.Map
	join       chan clientEvent
	leave      chan clientEvent
}

func (cm *channelManager) getChannel(channel string) (*Channel, error) {
	if value, ok := cm.channels.Load(channel); ok {
		return value.(*Channel), nil
	}
	value, err := cm.repository.FetchChannel(channel)
	if err != nil {
		cm.channels.Store(channel, value)
	}
	return value, err
}

func (cm *channelManager) Join(channel string, client *Client) {
	cm.join <- clientEvent{client, channel}
}

func (cm *channelManager) Leave(channel string, client *Client) {
	cm.leave <- clientEvent{client, channel}
}

func (cm *channelManager) Run() {
	for {
		select {
		case event := <-cm.join:
			channel, err := cm.getChannel(event.channel)
			if err != nil {
				log.Println("manager.Run:", err)
				// TODO: notify the client
				continue
			}
			channel.register <- event.client
		case event := <-cm.leave:
			channel, err := cm.getChannel(event.channel)
			if err != nil {
				log.Println("manager.Run:", err)
				// TODO: notify the client
				continue
			}
			channel.unregister <- event.client
		}
	}
}

func NewChannelManager(repository Repository) ChannelManager {
	return &channelManager{
		repository: repository,
		join:       make(chan clientEvent),
		leave:      make(chan clientEvent),
	}
}
