package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LPvdT/go-with-tests/application/internal/cli"
	"github.com/LPvdT/go-with-tests/application/internal/filesystem"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := filesystem.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	// HACK: Temporary dummy alerter, replace with real implementation if needed
	dummySpyAlerter := &cli.SpyBlindAlerter{}
	cli.NewCLI(store, os.Stdin, dummySpyAlerter).PlayPoker()
}
