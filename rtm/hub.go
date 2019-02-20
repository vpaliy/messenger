package rtm

import (
	"log"
	"sync"
)

type Hub struct {
	name        string
	subscribers map[*Client]bool
	broadcast   chan *ResponseMessage
	register    chan *Client
	unregister  chan *Client
	mutex       sync.Mutex
}

func NewHub(name string) *Hub {
	hub := Hub{
		name:        name,
		subscribers: make(map[*Client]bool),
		broadcast:   make(chan *ResponseMessage),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
	}
	go hub.Run()
	return &hub
}

func (h *Hub) HasClient(c *Client) bool {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	_, ok := h.subscribers[c]
	return ok
}

func (c *Hub) Run() {
	for {
		select {
		case client := <-c.register:
			// TODO: register in database
			client.Subscribe(&Subscription{
				c.name,
				c.broadcast,
			})
			c.subscribers[client] = true
		case client := <-c.unregister:
			// TODO: register in database
			client.Unsubscribe(c.name)
			if _, ok := c.subscribers[client]; ok {
				delete(c.subscribers, client)
			}
		case message := <-c.broadcast:
			// TODO: handle this better
			// TODO: put this change in the database
			log.Println("hub.Run: broadcasting")
			for sub, connected := range c.subscribers {
				if connected {
					go sub.JSON(message)
				}
			}
		}
	}
}
