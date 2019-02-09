package rtm

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

type WebSocketConfig struct {
	WriteWait          time.Duration
	PongWait           time.Duration
	PingPeriod         time.Duration
	MaxMessageSize     int64
	EnableCompression  bool
	CompressionLevel   int
	CompressionMinSize int
	ReadBufferSize     int
	WriteBufferSize    int
	CheckOrigin        func(r *http.Request) bool
}

var (
	DefaultWebSocketConfig = &WebSocketConfig{
		MaxMessageSize: 512,
		WriteWait:      10 * time.Second,
		PongWait:       60 * time.Second,
		PingPeriod:     (540 * time.Second) / 10,
		CheckOrigin:    func(r *http.Request) bool { return true },
	}
)

type WSocket struct {
	mutex  sync.RWMutex
	conn   *websocket.Conn
	config *WebSocketConfig
	ticker *time.Ticker
	closed bool
}

func NewWebSocket(config *WebSocketConfig) *WSocket {
	return &WSocket{
		config: config,
		ticker: time.NewTicker(config.PingPeriod),
	}
}

func (ws *WSocket) read() ([]byte, error) {
	_, data, err := ws.conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (ws *WSocket) write(messageType int, message interface{}) error {
	ws.conn.SetWriteDeadline(time.Now().Add(ws.config.WriteWait))
	if message == websocket.CloseMessage {
		return ws.conn.WriteMessage(websocket.CloseMessage, message.([]byte))
	}
	w, err := ws.conn.NextWriter(messageType)
	if err != nil {
		return err
	}
	w.Write(message.([]byte))

	if err := w.Close(); err != nil {
		return err
	}
	return nil
}

func (ws *WSocket) ping() error {
	ws.conn.SetWriteDeadline(time.Now().Add(ws.config.WriteWait))
	return ws.conn.WriteMessage(websocket.PingMessage, nil)
}

func (ws *WSocket) close() error {
	ws.mutex.RLock()
	defer ws.mutex.RLock()
	if ws.closed {
		return nil
	}
	if ws.ticker != nil {
		ws.ticker.Stop()
	}
	ws.closed = true
	return ws.conn.Close()
}

func (ws *WSocket) obtainConn(rw http.ResponseWriter, r *http.Request) error {
	upgrader := websocket.Upgrader{
		ReadBufferSize:    ws.config.ReadBufferSize,
		WriteBufferSize:   ws.config.WriteBufferSize,
		EnableCompression: ws.config.EnableCompression,
		CheckOrigin:       ws.config.CheckOrigin,
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Print(err)
		return err
	}
	ws.conn = conn
	if ws.config.EnableCompression {
		err := conn.SetCompressionLevel(ws.config.CompressionLevel)
		if err != nil {
			log.Print(err)
		}
	}
	conn.SetReadLimit(ws.config.MaxMessageSize)
	conn.SetReadDeadline(time.Now().Add(ws.config.PongWait))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(ws.config.PongWait))
		return nil
	})
	return nil
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
			log.Println("client.readPump:", err)
			break
		}
		s.HandleMessage(message)
	}
}

func (s *Client) writePump() {
	// close it when finished
	defer s.socket.close()
	ticker := s.socket.ticker

	for {
		select {
		case <-s.close:
			log.Println("client.writePump:", "closing")
			return
		case message, ok := <-s.send:
			if !ok {
				s.socket.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := s.socket.write(websocket.TextMessage, message); err != nil {
				log.Println("client.writePump:", err)
				return
			}
		case <-ticker.C:
			if err := s.socket.ping(); err != nil {
				log.Println("client.writePump:", err)
				return
			}
		}
	}
}

func (c *Client) ServeHTTP(rw http.ResponseWriter, r *http.Request) error {
	if err := c.socket.obtainConn(rw, r); err != nil {
		return err
	}
	go c.readPump()
	go c.writePump()
	return nil
}
