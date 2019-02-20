package rtm

import (
	"encoding/json"
	"log"
	"sync"
	"time"
)

type Client struct {
	socket        *WSocket
	uid           string
	token         string
	subscriptions map[string]*Subscription
	mutex         sync.Mutex
	send          chan interface{}
	manager       ChannelManager
	close         chan struct{}
}

type Subscription struct {
	channel   string
	broadcast chan<- *ResponseMessage
}

func NewClient(socket *WSocket, manager ChannelManager, token string) *Client {
	return &Client{
		socket:        socket,
		subscriptions: make(map[string]*Subscription),
		send:          make(chan interface{}),
		manager:       manager,
		token:         token,
	}
}

func (s *Client) Subscribe(sub *Subscription) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.subscriptions[sub.channel] = sub
}

func (c *Client) isSubscribed(channel string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_, ok := c.subscriptions[channel]
	return ok
}

func (s *Client) Unsubscribe(channel string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.subscriptions[channel]; ok {
		delete(s.subscriptions, channel)
	}
}

func (c *Client) HandleMessage(raw []byte) {
	request := new(WSEvent)
	if err := request.Unmarshal(raw); err != nil {
		log.Println("client.HandleMessage:", err)
		return
	}
	switch request.Event {
	case Join, Leave:
		var action *api.ChannelRequest
		if err := request.DecodeAction(action); err != nil {
			// TODO: send an error message
			return
		}
		if request.Event == Join {
			go c.Join(action)
		} else {
			go c.Leave(action)
		}
	case Hello:
		var action *ChannelRequest
		if err := request.DecodeAction(action); err != nil {
			// TODO: send an error message
			return
		}
		go c.Hello(action)
	case Send:
		var action *CreateMessageRequest
		if err := request.DecodeAction(action); err != nil {
			// TODO: send an error message
			return
		}
		go c.Send(action)
	case Typing:
		var action *ChannelRequest
		if err := request.DecodeAction(action); err != nil {
			// TODO: send an error message
			return
		}
		go c.Typing(action)
	case Load:
		var action *FetchMessagesRequest
		if err := request.DecodeAction(action); err != nil {
			// TODO: send an error message
			return
		}
		go c.Load(action)
	default:
		break
	}
}

func (c *Client) JSONError(err error) {
	// TODO: implement
}

func (c *Client) JSON(message *ResponseMessage) {
	serialized, err := json.Marshal(message)
	if err != nil {
		log.Println("client.JSON:", err)
		return
	}
	select {
	case c.send <- serialized:
	case <-time.After(time.Microsecond * 50):
		log.Println("client.JSON:", "Failed to send response")
		return
	}
}

func (c *Client) Join(req *ChannelRequest) {
	c.manager.Join(req.Channel, c)
}

func (c *Client) Leave(req *ChannelRequest) {
	c.manager.Leave(req.Channel, c)
}

func (c *Client) Send(req *CreateMessageRequest) {
	if sub, ok := c.subscriptions[req.Channel]; ok {
		response, err := c.dispatcher.Post(req, c)
		if err != nil {
			return
		}
		sub.broadcast <- response
	}
	// TODO: handle the case when the user is not subscribed to a chat
}

func (c *Client) Typing(action *ChannelRequest) {
	// TODO: handle this
}

func (c *Client) Load(action *FetchMessagesRequest) {
	// TODO: handle this
}

func (c *Client) Hello(action *HelloAction) {
	// TODO: handle this
}
