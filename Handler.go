package main

import (
	"encoding/json"
	"image/color"
	"net/http"
	"time"
)

var TICKER (*ColorTicker) = NewColorTicker(30 * time.Second)

type ColorResponse struct {
	Color         string
	Complementary [2]string
	Triadic       []string
	Tetradic      []string
}

type JsendResponse struct {
	Data   ColorResponse
	Status int
}

// Return a ColorResponse
func MakeColorResponse(c color.RGBA) (resp ColorResponse) {
	comp := Complementary(c)
	resp.Color = Hex(c)
	resp.Complementary = [2]string{
		Hex(comp[0]),
		Hex(comp[1]),
	}
	return
}

func ColorHandler(w http.ResponseWriter, r *http.Request) {
	var resp JsendResponse
	resp.Status = 200
	resp.Data = MakeColorResponse(TICKER.Color())

	js, err := json.Marshal(resp)

	if err != nil {
		w.Write([]byte("whatever"))
	} else {
		w.Write(js)
	}
}
