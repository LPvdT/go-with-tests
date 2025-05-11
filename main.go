package main

import (
	"log"
	"net/http"

	"github.com/LPvdT/go-with-tests/server"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
