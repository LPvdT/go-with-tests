// Continue here: https://quii.gitbook.io/learn-go-with-tests/build-an-application/io#didnt-we-just-break-some-rules-there-testing-private-things-no-interfaces

// Repository: https://github.com/LPvdT/go-with-tests/tree/feature/json-routing-embedding

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/LPvdT/go-with-tests/application/internal/common"
	"github.com/LPvdT/go-with-tests/application/internal/filesystem"
	"github.com/LPvdT/go-with-tests/application/server"
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
	store := &filesystem.FileSystemPlayerStore{
		Database: json.NewEncoder(&common.Tape{File: db}),
	}

	// Initialize the player server with the player store.
	server := server.NewPlayerServer(store)

	// Log the server address and start listening for HTTP requests.
	log.Printf("Starting server on http://%s", serverAddress)
	if err := http.ListenAndServe(serverAddress, server); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
