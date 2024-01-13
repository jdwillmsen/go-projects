package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Room struct {

	// clients holds all current clients in this room
	clients map[*Client]bool

	// join is a channel for clients wishing to join the room
	join chan *Client

	// leave is a channel for clients wishing to leave the room
	leave chan *Client

	// forward is a channel that holds incoming messages that should be forwarded to the other clients
	forward chan []byte
}

// newRoom creates a new chat room
func newRoom() *Room {
	return &Room{
		clients: make(map[*Client]bool),
		join:    make(chan *Client),
		leave:   make(chan *Client),
		forward: make(chan []byte),
	}
}

func (r *Room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.receive)
		case msg := <-r.forward:
			for client := range r.clients {
				client.receive <- msg
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP: ", err)
		return
	}
	client := &Client{
		socket:  socket,
		receive: make(chan []byte, messageBufferSize),
		room:    r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
