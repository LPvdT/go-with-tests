package cli_test

import (
	"strings"
	"testing"

	"github.com/LPvdT/go-with-tests/application/internal/cli"
	"github.com/LPvdT/go-with-tests/application/playertest"
)

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
