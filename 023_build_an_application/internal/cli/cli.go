package cli

import (
	"bufio"
	"io"
	"strings"

	"github.com/LPvdT/go-with-tests/application/internal/server"
)

type CLI struct {
	PlayerStore server.PlayerStore
	In          bufio.Scanner
}

func NewCLI(store server.PlayerStore, in io.Reader) *CLI {
	return &CLI{
		PlayerStore: store,
		In:          *bufio.NewScanner(in),
	}
}

func (cli *CLI) PlayPoker() {
	userInput := cli.readLine()
	cli.PlayerStore.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.In.Scan()
	return cli.In.Text()
}
