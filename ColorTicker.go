package main

import (
	"image/color"
	"sync"
	"time"
)

// Define a ColorTicker
type ColorTicker struct {
	color  color.RGBA
	ticker *time.Ticker
	mutex  *sync.Mutex
	Pool   *WebSocketConnectionPool
}

// Return a new ticker
func NewColorTicker(dur time.Duration) (s *ColorTicker) {

	s = &ColorTicker{
		RandomRGB(),
		time.NewTicker(dur),
		&sync.Mutex{},
		NewWebSocketConnectionPool(),
	}

	// Detach and start consumer ticker's channel
	go func() {
		for {
			s.mutex.Lock()
			s.color = RandomRGB()
			s.mutex.Unlock()

			<-s.ticker.C

			s.Pool.EmitAll(s.color)
		}
	}()

	return
}

// Stops ColorTicker
func (s *ColorTicker) Stop() {
	s.ticker.Stop()
}

// Returns an RGBA color
func (s *ColorTicker) Color() color.RGBA {
	var result color.RGBA

	s.mutex.Lock()
	result = s.color
	s.mutex.Unlock()

	return result
}
