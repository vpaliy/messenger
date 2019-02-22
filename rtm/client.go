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
	dispatcher    Dispatcher
	close         chan struct{}
}

type Subscription struct {
	channel   string
	broadcast chan<- *ResponseMessage
}

func NewClient(socket *WSocket, dispatcher Dispatcher, token string) *Client {
	return &Client{
		socket:        socket,
		subscriptions: make(map[string]*Subscription),
		send:          make(chan interface{}),
		dispatcher:    dispatcher,
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
		var action *ChannelRequest
		if err := request.DecodeAction(action); err != nil {
			c.JSONError(err)
			return
		}
		if request.Event == Join {
			go c.Join(action)
		} else {
			go c.Leave(action)
		}
	case Send:
		var action *CreateMessageRequest
		if err := request.DecodeAction(action); err != nil {
			c.JSONError(err)
			return
		}
		go c.Send(action)
	case Typing:
		var action *ChannelRequest
		if err := request.DecodeAction(action); err != nil {
			c.JSONError(err)
			return
		}
		go c.Typing(action)
	case Load:
		var action *FetchMessagesRequest
		if err := request.DecodeAction(action); err != nil {
			c.JSONError(err)
			return
		}
		go c.Load(action)
	default:
		// TODO: implement this
		break
	}
}

func (c *Client) JSONError(err error) {
	// TODO: implement
}

func (c *Client) JSON(message *ResponseMessage) {
	serialized, err := json.Marshal(message)
	if err != nil {
		c.JSONError(err)
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
	/* response, err := c.dispatcher.Leave(channelEvent{c, req.Channel})
	if err != nil {
		c.JSONError(err)
	} else {
		c.JSON(response)
	} */
}

func (c *Client) Leave(req *ChannelRequest) {
	/* response, err := c.dispatcher.Leave(channelEvent{c, req.Channel})
	if err != nil {
		c.JSONError(err)
	} else {
		c.JSON(response)
	} */
}

func (c *Client) Send(req *CreateMessageRequest) {
	if sub, ok := c.subscriptions[req.Channel]; ok {
		response, err := c.dispatcher.Post(postEvent{c.token, req})
		if err != nil {
			return
		}
		//sub.broadcast <- response
	}
	// TODO: handle the case when the user is not subscribed to a chat
}

func (c *Client) Typing(action *ChannelRequest) {
	// TODO: handle this
}

func (c *Client) Load(action *FetchMessagesRequest) {
	// TODO: handle this
}
