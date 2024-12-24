package main

import (
	"net/http"

	"github.com/BryanJ1ang/Dungeon-Crawler/internal/network"
)

func main() {
	// mux := http.NewServeMux()

	// mux.Handle("/", &homeHandler{})
	// mux.Handle("/start-game", &gameHandler{})

	// http.ListenAndServe(":8080", mux)

	network.StartGameServer("7777")
}

type gameHandler struct{}

type homeHandler struct{}

// Home handler ServerHTTP, serves the home page
func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}

// Game handler ServerHTTP, tells client to create a UDP connection for gameplay
func (h *gameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Start"))
	// send Game dungeon state as well
	// call go startGameServer("8081")
}
