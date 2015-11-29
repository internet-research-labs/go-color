package main

import (
	"github.com/gorilla/websocket"
	"image/color"
	"net/http"
)

// Websocket connection
type WebSocket struct {
	ws   *websocket.Conn
	send chan color.RGBA
	quit chan bool
}

// Global Upgrader
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
		send: make(chan color.RGBA),
		quit: make(chan bool),
		ws:   ws,
	}

	return &c, nil
}

func (s *WebSocket) Emit(message color.RGBA) {
	s.send <- message
}
