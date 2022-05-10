package ws

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// ClientManager is a websocket manager
type ClientManager struct {
	Clients    map[string]*Client
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

// Client is a websocket client
type Client struct {
	ID     string
	Socket *websocket.Conn
	Send   chan []byte
}

// Message is return msg
type Message struct {
	Type string      `json:"type,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// Manager defines a ws server manager
var Manager = ClientManager{
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[string]*Client),
}

// Start should be called on set up: go Manager.Start()
func (manager *ClientManager) Start() {
	for {
		select {
		case client := <-Manager.Register:
			log.Println("new ws connection", client.ID)
			Manager.Clients[client.ID] = client
			jsonMessage, _ := json.Marshal(&Message{Type: "message", Data: "connected"})
			client.Send <- jsonMessage
		case client := <-Manager.Unregister:
			log.Println("ws disconnected", client.ID)
			if _, ok := Manager.Clients[client.ID]; ok {
				jsonMessage, _ := json.Marshal(&Message{Type: "message", Data: "disconnected"})
				client.Send <- jsonMessage
				close(client.Send)
				delete(Manager.Clients, client.ID)
			}
		case message := <-Manager.Broadcast:
			for _, client := range Manager.Clients {
				client.Send <- message
			}
		}
	}
}
func creatId() string {
	return uuid.New().String()
}
func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		c.Socket.PongHandler()
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			return
		}
		log.Println("ws received", string(message))
		c.Send <- message
	}
}

func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				return
			}
			_ = c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

// WebsocketHandler handles websocket connect
func WebsocketHandler(c *gin.Context) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	client := &Client{
		ID:     creatId(),
		Socket: conn,
		Send:   make(chan []byte),
	}
	Manager.Register <- client
	go client.Read()
	go client.Write()
}
