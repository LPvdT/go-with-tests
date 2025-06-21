package cli

import (
	"time"

	"github.com/LPvdT/go-with-tests/application/internal/server"
)

type Game struct {
	alerter BlindAlerter
	store   server.PlayerStore
}

func NewGame(alerter BlindAlerter, store server.PlayerStore) *Game {
	return &Game{
		alerter: alerter,
		store:   store,
	}
}

func (p *Game) Start(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

func (p *Game) Finish(winner string) {
	p.store.RecordWin(winner)
}
