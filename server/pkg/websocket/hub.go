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

	// Inbound messages from the clients.
	// broadcast chan Message

	// Register requests from the clients.
	register chan *Client

	joinGame chan JoinInfo

	// Unregister requests from clients.
	unregister chan *Client
}

type JoinInfo struct {
	GameID string
	Client *Client
}

func newHub() *Hub {
	return &Hub{
		// broadcast:  make(chan Message),
		register:   make(chan *Client),
		joinGame:   make(chan JoinInfo),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
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
			// case message := <-h.broadcast:
			// 	//log.Printf("Broadcasting Message: %+v", message)
			// 	for client := range h.clients {
			// 		//fmt.Printf("Client Game ID: %v\n", client.gameID)
			// 		//fmt.Printf("Message Game ID: %v\n", message.GameID)
			// 		if message.GameID == client.gameID {
			// 			select {
			// 			case client.send <- message:
			// 			default:
			// 				close(client.send)
			// 				delete(h.clients, client)
			// 			}
			// 		}
			// 	}
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWsClient(hub *Hub, w http.ResponseWriter, r *http.Request) {

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

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}

func ServeWsServer(hub *Hub, w http.ResponseWriter, r *http.Request) {

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

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}
