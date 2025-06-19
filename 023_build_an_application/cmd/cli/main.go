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

	game := cli.NewGame(cli.BlindAlerterFunc(cli.StdOutAlerter), store)
	cli := cli.NewCLI(os.Stdin, os.Stdout, game)

	cli.PlayPoker()
}
