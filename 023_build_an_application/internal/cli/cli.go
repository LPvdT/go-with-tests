package cli

import (
	"bufio"
	"io"
	"strings"

	"github.com/LPvdT/go-with-tests/application/internal/server"
)

type CLI struct {
	PlayerStore server.PlayerStore
	In          io.Reader
}

func (cli *CLI) PlayPoker() {
	reader := bufio.NewScanner(cli.In)
	reader.Scan()
	cli.PlayerStore.RecordWin(extractWinner(reader.Text()))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
