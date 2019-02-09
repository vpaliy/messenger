package rtm

import "sync"

type Repository interface {
	FetchChannel(channel string) (*Channel, error)
}

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
	repository *Repository
	channels   sync.Map
	join       chan clientEvent
	leave      chan clientEvent
}

func (cm *channelManager) hasChannel(channel string) bool {
	res := cm.getChannel(channel)
	return res != nil
}

func (cm *channelManager) getChannel(channel string) *Channel {
	if value, ok := cm.channels.Load(channel); ok {
		return value.(*Channel)
	}
	return nil
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
		case <-cm.join:
			// TODO:
		case <-cm.leave:
			// TODO:
		}
	}
}

func NewChannelManager(repository *Repository) ChannelManager {
	return &channelManager{
		repository: repository,
		join:       make(chan clientEvent),
		leave:      make(chan clientEvent),
	}
}
