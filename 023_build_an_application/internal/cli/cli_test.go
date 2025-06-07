package cli_test

import (
	"strings"
	"testing"
	"time"

	"github.com/LPvdT/go-with-tests/application/internal/cli"
	"github.com/LPvdT/go-with-tests/application/playertest"
)

type SpyBlindAlerter struct {
	Alerts []struct {
		scheduledAt time.Duration
		amount      int
	}
}

// type failOnEndReader struct {
// 	t   *testing.T
// 	rdr io.Reader
// }

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.Alerts = append(s.Alerts, struct {
		scheduledAt time.Duration
		amount      int
	}{duration, amount})
}

var dummySpyAlerter = &SpyBlindAlerter{}

func TestCLI(t *testing.T) {
	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &playertest.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := cli.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		if len(blindAlerter.Alerts) != 1 {
			t.Fatal("expected a blind alert to be scheduled")
		}
	})

	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader(("Chris wins\n"))
		playerStore := &playertest.StubPlayerStore{}

		cli := cli.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		playertest.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &playertest.StubPlayerStore{}

		cli := cli.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		playertest.AssertPlayerWin(t, playerStore, "Cleo")
	})

	// t.Run("do not read beyond the first newline", func(t *testing.T) {
	// 	in := failOnEndReader{
	// 		t,
	// 		strings.NewReader("Chris winds\n hello there"),
	// 	}
	// 	playerStore := &playertest.StubPlayerStore{}

	// 	cli := cli.NewCLI(playerStore, in, dummySpyAlerter)
	// 	cli.PlayPoker()
	// })
}
