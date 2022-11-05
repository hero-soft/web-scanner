package websocket

import (
	"bytes"
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
	trunkrecorder "github.com/hero-soft/web-scanner/pkg/trunk-recorder"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 4096
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	clientID string

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan string
}

func (c *Client) recorderReadPump() {
	defer func() {
		log.Println("Client unregistered: ", c.clientID)
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	c.conn.SetPingHandler(func(string) error { return nil })

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		//unmarshall message envelope
		var messageEnvelope trunkrecorder.MessageEnvelope
		err = json.Unmarshal(message, &messageEnvelope)

		if err != nil {
			log.Println("Error unmarshalling message envelope: ", err)
			break
		}

		switch messageEnvelope.MessageType {
		case "rates":
			//unmarshall rates message
		case "call_end":
			log.Printf("Call end message: %v", string(message))
		default:
			// log.Println("Got message from trunk-recorder")
			// log.Printf("MESSAGE %s", messageEnvelope.MessageType, string(message))
		}

		// var messageJSON Message
		// err = json.Unmarshal(bytes.TrimSpace(bytes.Replace(message, newline, space, -1)), &messageJSON)

		// if err != nil {
		// 	log.Printf("Could not unmarshall message JSON: %v", err)
		// 	continue
		// }

		c.hub.broadcast <- string(message)
	}
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) clientReadPump() {
	defer func() {
		log.Println("Client unregistered: ", c.clientID)
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	c.conn.SetPingHandler(func(string) error { return nil })

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		log.Printf("MESSAGE from %s: %v", c.clientID, string(message))

	}
}

// clientWritePump pumps messages from the hub to the websocket connection.
//
// A goroutine running clientWritePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) clientWritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			// // messageBytes, err := json.Marshal(message)
			messageBytes := bytes.NewBufferString(message).Bytes()

			// if err != nil {
			// 	log.Printf("Could not marshall message: %v", err)
			// }

			w.Write(messageBytes)

			if err := w.Close(); err != nil {
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
