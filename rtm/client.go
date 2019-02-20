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
	handler       EventHandler
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
	request := new(ActionRequest)
	if err := json.Unmarshal(raw, request); err != nil {
		log.Println("client.HandleMessage:", err)
		return
	}
	log.Println("client.HandleMessage:", request)
	switch request.Event {
	case Join, Leave:
		action, err := new(JoinAction).Decode(request)
		if err != nil {
			log.Println("client.HandleMessage:", err)
			return
		}
		if request.Event == Join {
			go c.Join(action)
		} else {
			go c.Leave(action)
		}
	case Hello:
		action, err := new(HelloAction).Decode(request)
		if err != nil {
			log.Println("client.HandleMessage:", err)
			return
		}
		go c.Hello(action)
	case Send:
		action, err := new(MessageAction).Decode(request)
		if err != nil {
			log.Println("client.HandleMessage:", err)
			return
		}
		go c.Send(action)
	case Typing:
		action, err := new(TypingAction).Decode(request)
		if err != nil {
			log.Println("client.HandleMessage:", err)
			return
		}
		go c.Typing(action)
	case Load:
		action, err := new(LoadRequest).Decode(request)
		if err != nil {
			log.Println("client.HandleMessage:", err)
			return
		}
		go c.Load(action)
	default:
		break
	}
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

func (c *Client) Join(action *JoinAction) {
	c.manager.Join(action.Channel, c)
}

func (c *Client) Leave(action *JoinAction) {
	c.manager.Leave(action.Channel, c)
}

func (c *Client) Send(action *MessageAction) {
	if err := c.manager.Post(action); err != nil {
		// TODO: send an error message here
		return
	}
	if sub, ok := c.subscriptions[action.Channel]; ok {
		response := &ResponseMessage{Send, action.Content}
		sub.broadcast <- response
		return
	}
	// TODO: handle the case when the user is not subscribed to a chat
}

func (c *Client) Typing(action *TypingAction) {
	// TODO: handle this
}

func (c *Client) Load(action *LoadRequest) {
	// TODO: handle this
}

func (c *Client) Hello(action *HelloAction) {
	// TODO: handle this
}
