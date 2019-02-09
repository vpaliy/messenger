package rtm

type Channel struct {
	name        string
	subscribers map[*Client]bool
	broadcast   chan *ResponseMessage
	register    chan *Client
	unregister  chan *Client
}

func NewChannel(name string) *Channel {
	channel := Channel{
		name:        name,
		subscribers: make(map[*Client]bool),
		broadcast:   make(chan *ResponseMessage),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
	}
	go channel.Run()
	return &channel
}

func (c *Channel) Run() {
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
			for sub, connected := range c.subscribers {
				if connected {
					go sub.JSON(message)
				}
			}
		}
	}
}
