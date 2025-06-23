package cli_test

// BUG: Tests don't work. Just fix later and continue with the book.

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/LPvdT/go-with-tests/application/common"
	mod_cli "github.com/LPvdT/go-with-tests/application/internal/cli"
)

type ScheduledAlert struct {
	scheduledAt time.Duration
	amount      int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.scheduledAt)
}

type SpyBlindAlerter struct {
	alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, ScheduledAlert{at, amount})
}

var (
	dummyBlindAlerter = &SpyBlindAlerter{}
	dummyPlayerStore  = &common.StubPlayerStore{}
	dummyStdOut       = &bytes.Buffer{}
	// dummyStdIn        = &bytes.Buffer{}
)

func TestCLI(t *testing.T) {
	t.Run("it schedules the printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		blindAlerter := &SpyBlindAlerter{}
		game := mod_cli.NewGame(blindAlerter, dummyPlayerStore)

		cli := mod_cli.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		cases := []ScheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})

	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		game := mod_cli.NewGame(dummyBlindAlerter, dummyPlayerStore)

		cli := mod_cli.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		common.AssertPlayerWin(t, dummyPlayerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		game := mod_cli.NewGame(dummyBlindAlerter, dummyPlayerStore)

		cli := mod_cli.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		common.AssertPlayerWin(t, dummyPlayerStore, "Cleo")
	})

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		blindAlerter := &SpyBlindAlerter{}
		game := mod_cli.NewGame(dummyBlindAlerter, dummyPlayerStore)

		cli := mod_cli.NewCLI(in, stdout, game)
		cli.PlayPoker()

		got := stdout.String()
		want := mod_cli.PlayerPrompt

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		cases := []ScheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})
}

func assertScheduledAlert(t testing.TB, got, want ScheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
