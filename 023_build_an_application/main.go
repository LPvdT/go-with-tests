// Continue here: https://quii.gitbook.io/learn-go-with-tests/build-an-application/io#didnt-we-just-break-some-rules-there-testing-private-things-no-interfaces

// Repository: https://github.com/LPvdT/go-with-tests/tree/feature/json-routing-embedding

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/LPvdT/go-with-tests/application/common"
	"github.com/LPvdT/go-with-tests/application/internal/filesystem"
	"github.com/LPvdT/go-with-tests/application/server"
)

var (
	port          string = "5000"
	serverAddress string = strings.Join([]string{"localhost", port}, ":")
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0o666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store := &filesystem.FileSystemPlayerStore{
		Database: json.NewEncoder(&common.Tape{File: db}),
	}

	server := server.NewPlayerServer(store)

	log.Printf("Starting server on http://%s", serverAddress)
	if err := http.ListenAndServe(serverAddress, server); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
