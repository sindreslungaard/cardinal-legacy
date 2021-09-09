package network

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 80 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = 40 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 4096
)

type WebsocketClient struct {
	conn *websocket.Conn
	send chan interface{}

	disposed bool
}

func NewWebsocketClient(conn *websocket.Conn) *WebsocketClient {
	return &WebsocketClient{
		conn: conn,
		send: make(chan interface{}),

		disposed: false,
	}
}

func BeginReading(c *WebsocketClient) {
	defer func() {
		Dispose(c)
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {

		_, data, err := c.conn.ReadMessage()

		if err != nil {
			return
		}

		ParseMessage(c, data)

	}
}

func BeginWriting(c *WebsocketClient) {
	defer func() {
		Dispose(c)
	}()

	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for {
		select {
		case message, ok := <-c.send:

			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := c.conn.WriteJSON(message)

			if err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}

		}
	}
}

func Send(c *WebsocketClient, data interface{}) {
	if c.disposed {
		return
	}

	c.send <- data
}

func Dispose(c *WebsocketClient) {
	if c.disposed {
		return
	}

	c.disposed = true

	c.conn.Close()
	close(c.send)
}

func ParseMessage(c *WebsocketClient, data []byte) {

}
