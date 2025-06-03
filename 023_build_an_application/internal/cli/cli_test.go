package cli_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/LPvdT/go-with-tests/application/internal/cli"
	"github.com/LPvdT/go-with-tests/application/playertest"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader(("Chris wins\n"))
		playerStore := &playertest.StubPlayerStore{}

		cli := &cli.CLI{PlayerStore: playerStore, In: *bufio.NewScanner(in)}
		cli.PlayPoker()

		playertest.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader(("Cleo wins\n"))
		playerStore := &playertest.StubPlayerStore{}

		cli := &cli.CLI{PlayerStore: playerStore, In: *bufio.NewScanner(in)}
		cli.PlayPoker()

		playertest.AssertPlayerWin(t, playerStore, "Cleo")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &playertest.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := cli.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		if len(blindAlerter.alerts) != 1 {
			t.Fatal("expected a blind alert to be scheduled")
		}
	})
}

}
