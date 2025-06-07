package cli

import (
	"bufio"
	"io"
	"strings"
	"time"

	"github.com/LPvdT/go-with-tests/application/internal/server"
)

type CLI struct {
	PlayerStore server.PlayerStore
	In          bufio.Scanner
	Alerter     BlindAlerter
}

func NewCLI(store server.PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{
		PlayerStore: store,
		In:          *bufio.NewScanner(in),
		Alerter:     alerter,
	}
}

func (cli *CLI) PlayPoker() {
	cli.Alerter.ScheduleAlertAt(10*time.Second, 100)
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

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}
