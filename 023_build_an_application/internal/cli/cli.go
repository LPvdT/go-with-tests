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
