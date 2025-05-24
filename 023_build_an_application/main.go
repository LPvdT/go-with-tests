// Repository: https://github.com/LPvdT/go-with-tests/tree/feature/json-routing-embedding

package main

import (
	"log"
	"net/http"

	"github.com/LPvdT/go-with-tests/application/server"
	"github.com/LPvdT/go-with-tests/application/store"
)

func main() {
	store := store.NewInMemoryPlayerStore()
	server := server.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
