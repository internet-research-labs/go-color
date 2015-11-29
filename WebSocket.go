package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// Websocket connection
type WebSocket struct {
	ws   *websocket.Conn
	send chan []byte
	quit chan bool
}

var UPGRADER = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Return a Wrapper Websocket Connection From an HTTP Request
func NewWebSocket(w http.ResponseWriter, r *http.Request) (*WebSocket, error) {
	ws, err := UPGRADER.Upgrade(w, r, nil)

	if err != nil {
		return nil, err
	}

	c := WebSocket{
		send: make(chan []byte),
		quit: make(chan bool),
		ws:   ws,
	}

	return &c, nil
}

func (s *WebSocket) Emit(message []byte) {
	s.send <- message
}

// Create a Persistent Websocket Connection
func SocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := NewWebSocket(w, r)

	if err != nil {
		return
	}

	log.Print("FUCKKKK")

	go func() {
		for {
			select {
			case <-conn.send:
				conn.ws.WriteMessage(websocket.TextMessage, []byte("sick"))
			case <-conn.quit:
				break
			}
		}
	}()
}
