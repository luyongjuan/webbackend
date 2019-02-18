package pusher

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = pongWait * 9 / 10
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Clienter interface {
	Send(msg []byte)
}

type Client struct {
	hub     *Hub
	conn    *websocket.Conn
	sendCh  chan []byte
	closeCh chan struct{}
}

func (c *Client) Send(msg []byte) {
	select {
	case <-c.closeCh:
		return
	case c.sendCh <- msg:
	}
}

func (c *Client) read() {
	defer func() {
		c.hub.unregister <- c
		close(c.closeCh)
		c.conn.Close()
	}()

	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("err: %v", err)
			}
			break
		}

		c.hub.handleMsg(message, c)
	}
}

func (c *Client) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.sendCh:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := c.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("write message err: %v", err)
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

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{
		hub:     hub,
		conn:    conn,
		sendCh:  make(chan []byte),
		closeCh: make(chan struct{}),
	}
	client.hub.register <- client

	go client.write()
	go client.read()
}
