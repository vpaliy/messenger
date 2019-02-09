package rtm

import (
	"encoding/json"
	"sync"
	"time"
)

type Client struct {
	socket        *WSocket
	uid           string
	user          string
	subscriptions map[string]*Subscription
	mutex         sync.Mutex
	send          chan interface{}
	manager       ChannelManager
	handler       EventHandler
}

type Subscription struct {
	channel   string
	broadcast chan<- *ResponseMessage
}

func NewClient(socket *WSocket, manager ChannelManager, handler EventHandler) *Client {
	return &Client{
		socket:        socket,
		subscriptions: make(map[string]*Subscription),
		send:          make(chan interface{}),
		manager:       manager,
		handler:       handler,
	}
}

func (s *Client) readPump() {
	defer func() {
		// TODO: unregister from hub
		s.socket.close()
	}()
	// listen
	for {
		message, err := s.socket.read()
		if err != nil {
			break
		}
		s.HandleMessage(message)
	}
}

func (s *Client) writePump() {
	config := s.socket.config
	ticker := time.NewTicker(config.pingPeriod)
	// close it when finished
	defer func() {
		ticker.Stop()
		s.socket.close()
	}()

	for {
		select {
		case message, ok := <-s.send:
			if !ok {
				// TODO: write close message
				return
			}
			if err := s.socket.write(message); err != nil {
				return
			}
		case <-ticker.C:
			if err := s.socket.ping(); err != nil {
				return
			}
		}
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
		return
	}
	switch request.Event {
	case Join, Leave:
		action, err := new(JoinAction).Decode(request)
		if err != nil {
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
			return
		}
		go c.Hello(action)
	case Send:
		action, err := new(MessageAction).Decode(request)
		if err != nil {
			return
		}
		go c.Send(action)
	case Typing:
		action, err := new(TypingAction).Decode(request)
		if err != nil {
			return
		}
		go c.Typing(action)
	case Load:
		action, err := new(LoadRequest).Decode(request)
		if err != nil {
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
		// TODO: log the error
	}
	select {
	case c.send <- serialized:
	case <-time.After(time.Microsecond * 50):
		// TODO: log that we were unable to send this
	}
}

func (c *Client) Join(action *JoinAction) {
	c.manager.Join(action.Channel, c)
}

func (c *Client) Leave(action *JoinAction) {
	c.manager.Leave(action.Channel, c)
}

func (c *Client) Send(action *MessageAction) {
	// TODO: handle this
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
