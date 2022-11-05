package websocket

import (
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	recorder *Client

	// Messages going to all clients
	broadcast chan string

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan string, 256),
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

			log.Println("Client registered: ", client.clientID)

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			//log.Printf("Broadcasting Message: %+v", message)
			for client := range h.clients {
				//fmt.Printf("Client Game ID: %v\n", client.gameID)
				//fmt.Printf("Message Game ID: %v\n", message.GameID)

				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}

			}
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
		log.Println(err)
		return
	}

	client := &Client{hub: hub, conn: conn, send: make(chan string, 256)}
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
		log.Println(err)
		return
	}

	// _, m, _ := conn.ReadMessage()
	// log.Println(string(m))

	client := &Client{hub: hub, conn: conn}
	hub.recorder = client

	go client.recorderReadPump()
}
