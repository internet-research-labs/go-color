package main

import (
	"image/color"
	"time"
)

type WebSocketConnectionPool struct {
	conns  [](*WebSocket)
	ticker *time.Ticker
}

// Return a New Websocket Connection Pool
func NewWebSocketConnectionPool() *WebSocketConnectionPool {

	// Create Pool of Socket Connections
	s := &WebSocketConnectionPool{
		nil,
		time.NewTicker(1 * time.Second),
	}

	return s
}

// Add a WebSocket to the Connection Pool
func (s *WebSocketConnectionPool) Add(ws *WebSocket) {
	s.conns = append(s.conns, ws)
}

// Remove a WebSocket from the Connection Pool
func (s *WebSocketConnectionPool) Remove(ws *WebSocket) {
	// TODO: Implement this when you learn how
}

// Remove every closed connection
func (s *WebSocketConnectionPool) RemoveClosed() {
	// TODO: Implement this when you learn how
}

// Emit to All Connections in Pool
func (s *WebSocketConnectionPool) EmitAll(message color.RGBA) {
	for _, v := range s.conns {
		v.Emit(message)
	}
}
