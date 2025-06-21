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
	"github.com/goforj/godump"
	"github.com/lmittmann/tint"
)

const (
	dbFileName   = "game.db.json"
	htmlDumpName = "game_dump.html"
	printDump    = false // Set to true to print the game state dump to the console
	htmlDump     = true  // Set to true to write the game state dump to an HTML file
)

var dumpState = fmt.Sprintf("no-dump-game-state%v", func() string {
	if !htmlDump {
		return "-html"
	}
	return ""
}())

func main() {
	logger := slog.Default()

	logger.Debug("Loading player store", "filename", dbFileName)
	store, close, err := filesystem.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	logger.Info("Let's play poker")
	logger.Info(
		"Type <name> wins to record a win",
		"<name>", "Teun", "<name>", "Teun wins",
	)

	game := cli.NewGame(cli.BlindAlerterFunc(cli.StdOutAlerter), store)
	cli := cli.NewCLI(os.Stdin, os.Stdout, game)

	if printDump {
		logger.Warn("Dumping game state", "sink", "stdout")
		godump.Dump(game)
	} else {
		logger.Error("Not dumping game state to console",
			tint.Err(errors.New(dumpState)), "sink", "stdout",
		)
	}

	if htmlDump {
		logger.Warn("Dumping game state to HTML file...", "filename", htmlDumpName)

		html := godump.DumpHTML(game)

		f, _ := os.OpenFile(htmlDumpName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
		defer f.Close()

		_, err = f.Write([]byte(html))
		if err != nil {
			panic(err)
		}
	} else {
		logger.Error("Not dumping game state to HTML file",
			tint.Err(errors.New(dumpState)), "filename", htmlDumpName,
		)
	}

	cli.PlayPoker()
}
