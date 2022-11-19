package websocket

import (
	"net/http"

	"github.com/hero-soft/web-scanner/pkg/call"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

// hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	recorder *Client

	logger *zap.SugaredLogger

	// Messages going to all clients
	broadcast chan Message

	Send chan SendTo

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

type SendTo struct {
	To      string  `json:"to"`
	Message Message `json:"message"`
}

type Message struct {
	Type  string      `json:"type,omitempty"`
	Call  call.Call   `json:"call,omitempty"`
	Calls []call.Call `json:"calls"`
}

func NewHub(logger *zap.SugaredLogger) *Hub {
	return &Hub{
		logger:     logger,
		broadcast:  make(chan Message, 256),
		Send:       make(chan SendTo, 256),
		register:   make(chan *Client, 256),
		unregister: make(chan *Client, 256),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true

			uuid := uuid.NewV4()
			client.clientID = uuid.String()

			h.logger.Infof("Client registered: %s", client.clientID)

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {

				select {
				case client.send <- message:

				default:
					close(client.send)
					delete(h.clients, client)
				}

			}
		case sendTo := <-h.Send:
			h.broadcast <- sendTo.Message
		}
	}
}

// serveWs handles websocket requests from the peer.
func (hub *Hub) ServeWsClient(w http.ResponseWriter, r *http.Request) {

	// if !localSettings.Production {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	// }

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		hub.logger.Error(err)
		return
	}

	client := &Client{logger: hub.logger, hub: hub, conn: conn, send: make(chan Message, 256)}
	client.hub.register <- client

	go client.clientWritePump()
	go client.clientReadPump()
}

func (hub *Hub) ServeWsRecorder(w http.ResponseWriter, r *http.Request) {

	// if !localSettings.Production {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	// }

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		hub.logger.Error(err)
		return
	}

	client := &Client{logger: hub.logger, hub: hub, conn: conn}
	hub.recorder = client

	go client.recorderReadPump()
}
