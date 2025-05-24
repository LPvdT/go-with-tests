// Continue here: https://quii.gitbook.io/learn-go-with-tests/build-an-application/json#write-the-test-first-4
// Repository: https://github.com/LPvdT/go-with-tests/tree/feature/json-routing-embedding

package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/LPvdT/go-with-tests/application/server"
	"github.com/LPvdT/go-with-tests/application/store"
)

var (
	port          string = "5000"
	serverAddress string = strings.Join([]string{"localhost", port}, ":")
)

func main() {
	store := store.NewInMemoryPlayerStore()
	server := server.NewPlayerServer(store)

	log.Printf("Starting server on http://%s", serverAddress)
	if err := http.ListenAndServe(serverAddress, server); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
