package main

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/LPvdT/go-with-tests/application/internal/cli"
	"github.com/LPvdT/go-with-tests/application/internal/filesystem"
	_ "github.com/LPvdT/go-with-tests/application/internal/logging"
	"github.com/lmittmann/tint"
)

const dbFileName = "game.db.json"

func main() {
	logger := slog.Default()

	logger.Debug("Loading player store:", "filename", dbFileName)
	logger.Info("Loading player store:", "filename", dbFileName)
	logger.Warn("Loading player store:", "filename", dbFileName)
	logger.Error("Loading player store:", tint.Err(errors.New("teun")), "filename", dbFileName)
	store, close, err := filesystem.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	game := cli.NewGame(cli.BlindAlerterFunc(cli.StdOutAlerter), store)
	cli := cli.NewCLI(os.Stdin, os.Stdout, game)

	cli.PlayPoker()
}
