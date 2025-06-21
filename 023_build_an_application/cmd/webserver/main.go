package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/LPvdT/go-with-tests/application/internal/filesystem"
	_ "github.com/LPvdT/go-with-tests/application/internal/logging"
	"github.com/LPvdT/go-with-tests/application/internal/server"
	"github.com/lmittmann/tint"
)

var (
	port          string = "5000"
	address       string = "localhost"
	serverAddress string = strings.Join([]string{address, port}, ":")
	dbFileName    string = "game.db.json"
)

func main() {
	logger := slog.Default()

	// Create a file system player store from the specified database file.
	store, close, err := filesystem.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		logger.Error(
			"Could not create player store from file",
			tint.Err(err), "filename", dbFileName,
		)
		os.Exit(1)
	}
	defer close()

	// Initialize the player server with the player store.
	server := server.NewPlayerServer(store)

	// Log the server address and start listening for HTTP requests.
	logger.Info(
		"Starting server on",
		"address", fmt.Sprintf("http://%s:%s", address, port),
		"league-endpoint", fmt.Sprintf("http://%s:%s/league", address, port),
	)
	if err := http.ListenAndServe(serverAddress, server); err != nil {
		logger.Error("Could not start server:", tint.Err(err), "address", serverAddress)
		os.Exit(1)
	}
}
