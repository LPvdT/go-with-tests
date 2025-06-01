package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/LPvdT/go-with-tests/application/internal/cli"
	"github.com/LPvdT/go-with-tests/application/internal/filesystem"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker!")
	fmt.Println("Type {Name} wins to record a win for {Name}.")

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0o666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := filesystem.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	game := cli.CLI{PlayerStore: store, In: *bufio.NewScanner(os.Stdin)}
	game.PlayPoker()
}
