package main

import (
	"log"
	"net/http"

	"github.com/LPvdT/go-with-tests/application/server"
)

func main() {
	server := &server.PlayerServer{}
	log.Fatal(http.ListenAndServe(":5000", server))
}
