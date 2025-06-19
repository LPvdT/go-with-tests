package cli

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/LPvdT/go-with-tests/application/internal/server"
)

const PlayerPrompt = "Please enter the number of players: "

type CLI struct {
	In   bufio.Scanner
	Out  io.Writer
	Game *Game
}

func NewCLI(store server.PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{
		In:  *bufio.NewScanner(in),
		Out: out,
		Game: &Game{
			alerter: alerter,
			store:   store,
		},
	}
}

func (cli *CLI) PlayPoker() {
	_, err := fmt.Fprint(cli.Out, PlayerPrompt)
	if err != nil {
		panic(err)
	}

	numberOfPlayersInput := cli.readLine()
	numberOfPlayers, _ := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

	cli.Game.Start(numberOfPlayers)

	winnerInput := cli.readLine()
	winner := extractWinner(winnerInput)

	cli.Game.Finish(winner)
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins\n", "", 1)
}

func (cli *CLI) readLine() string {
	cli.In.Scan()
	return cli.In.Text()
}
