package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/LPvdT/go-with-tests/application/internal/filesystem"
	"github.com/LPvdT/go-with-tests/application/internal/server"
)

var (
	port          string = "5000"
	address       string = "localhost"
	serverAddress string = strings.Join([]string{address, port}, ":")
)

const dbFileName string = "game.db.json"

func main() {
	// Create a file system player store from the specified database file.
	store, close, err := filesystem.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	// Initialize the player server with the player store.
	server := server.NewPlayerServer(store)

	// Log the server address and start listening for HTTP requests.
	log.Printf("Starting server on http://%s", serverAddress)
	if err := http.ListenAndServe(serverAddress, server); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
