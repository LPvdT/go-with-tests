package main

import (
	"log"
	"net/http"
	"os"
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
	// Open the database file with read and write permissions, creating it if it doesn't exist.
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0o666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	// Create a new file system player store with the database.
	store, err := filesystem.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store: %v", err)
	}

	// Initialize the player server with the player store.
	server := server.NewPlayerServer(store)

	// Log the server address and start listening for HTTP requests.
	log.Printf("Starting server on http://%s", serverAddress)
	if err := http.ListenAndServe(serverAddress, server); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
