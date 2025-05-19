package main

import (
	"log"
	"net/http"

	"github.com/LPvdT/go-with-tests/application/server"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &server.PlayerServer{}
	log.Fatal(http.ListenAndServe(":5000", server))
}
