package cli_test

import (
	"strings"
	"testing"

	"github.com/LPvdT/go-with-tests/application/internal/cli"
	"github.com/LPvdT/go-with-tests/application/playertest"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader(("Chris wins\n"))
		playerStore := &playertest.StubPlayerStore{}

		cli := &cli.CLI{PlayerStore: playerStore, In: in}
		cli.PlayPoker()

		playertest.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader(("Cleo wins\n"))
		playerStore := &playertest.StubPlayerStore{}

		cli := &cli.CLI{PlayerStore: playerStore, In: in}
		cli.PlayPoker()

		playertest.AssertPlayerWin(t, playerStore, "Cleo")
	})
}
