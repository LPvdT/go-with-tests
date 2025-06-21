package cli

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const PlayerPrompt = "Please enter the number of players: "

type CLI struct {
	In   *bufio.Scanner
	Out  io.Writer
	Game *Game
}

func NewCLI(in io.Reader, out io.Writer, game *Game) *CLI {
	return &CLI{
		In:   bufio.NewScanner(in),
		Out:  out,
		Game: game,
	}
}

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.Out, PlayerPrompt)

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
