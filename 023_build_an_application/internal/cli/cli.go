package cli

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/LPvdT/go-with-tests/application/internal/server"
)

const PlayerPrompt = "Please enter the number of players: "

type CLI struct {
	PlayerStore server.PlayerStore
	In          bufio.Scanner
	Out         io.Writer
	Alerter     BlindAlerter
}

func NewCLI(store server.PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{
		PlayerStore: store,
		In:          *bufio.NewScanner(in),
		Out:         out,
		Alerter:     alerter,
	}
}

func (cli *CLI) PlayPoker() {
	_, err := fmt.Fprint(cli.Out, PlayerPrompt)
	if err != nil {
		panic(err)
	}

	cli.scheduleBlindAlerts()
	userInput := cli.readLine()
	cli.PlayerStore.RecordWin(extractWinner(userInput))
}

func (cli *CLI) scheduleBlindAlerts() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		cli.Alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + 10*time.Minute
	}
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.In.Scan()
	return cli.In.Text()
}
