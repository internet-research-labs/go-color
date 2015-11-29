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

	fmt.Println("Swag")

	r := mux.NewRouter()
	r.HandleFunc("/", ColorHandler)

	http.ListenAndServe(":8080", r)
}
