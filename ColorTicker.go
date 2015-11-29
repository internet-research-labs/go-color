package main

import (
	"fmt"
	"image/color"
	"time"
)

// Define a ColorTicker
type ColorTicker struct {
	color  color.RGBA
	ticker *time.Ticker
}

// Return a new ticker
func NewColorTicker(dur time.Duration) (s *ColorTicker) {

	s = &ColorTicker{
		RandomRGB(),
		time.NewTicker(dur),
	}

	// Detach and start consumer ticker's channel
	go func() {
		for {
			fmt.Println(s.color)
			s.color = RandomRGB()
			<-s.ticker.C
		}
	}()

	return
}

// Stops ColorTicker
func (s ColorTicker) Stop() {
	s.ticker.Stop()
}

// Returns an RGBA color
func (s ColorTicker) Color() color.RGBA {
	fmt.Println(s.color)
	return s.color
}
