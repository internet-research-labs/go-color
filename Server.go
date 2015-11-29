package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	// Seed
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("Starting server...")

	r := mux.NewRouter()
	r.HandleFunc("/request", HttpHandler)
	r.HandleFunc("/socket", SocketHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./views/")))

	http.ListenAndServe(":8080", r)
}
