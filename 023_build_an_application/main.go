package main

import (
	"log"
	"net/http"

	"github.com/LPvdT/go-with-tests/build_an_application/server"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
