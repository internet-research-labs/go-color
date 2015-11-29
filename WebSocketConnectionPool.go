package main

type WebSocketConnectionPool struct {
	conns []WebSocket
}

func NewWebSocketConnectionPool() *WebSocketConnectionPool {
	return &WebSocketConnectionPool{nil}
}

func (s *WebSocketConnectionPool) Add(ws *WebSocket) {
}

func (s *WebSocketConnectionPool) Remove(ws *WebSocket) {
}

// Remove every closed connection
func (s *WebSocketConnectionPool) RemoveClosed() {
	// TODO: Implement this when you learn how
}

// Emit to All Connections in Pool
func (s *WebSocketConnectionPool) EmitAll(message []byte) {
	for _, v := range s.conns {
		v.Emit(message)
	}
}
