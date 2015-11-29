package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"image/color"
	"net/http"
	"time"
)

var TICKER (*ColorTicker) = NewColorTicker(1 * time.Second)

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

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := NewWebSocket(w, r)

	if err != nil {
		return
	}

	TICKER.Pool.Add(conn)

	go func() {
		for {

			select {
			case v := <-conn.send:
				fmt.Println(v)

				var resp JsendResponse
				resp.Status = 200
				resp.Data = MakeColorResponse(v)

				js, jsErr := json.Marshal(resp)

				if jsErr != nil {
					js = []byte("undefined")
				}

				conn.ws.WriteMessage(websocket.TextMessage, js)
			case <-conn.quit:
				break
			}
		}
	}()
}
