package main

import (
	"strings"
	"testing"

	"github.com/LPvdT/go-with-tests/application/internal/cli"
	"github.com/LPvdT/go-with-tests/application/internal/common"
)

func TestCLI(t *testing.T) {
	in := strings.NewReader(("Cleo wins\n"))
	playerStore := &common.StubPlayerStore{}

	cli := &cli.CLI{PlayerStore: playerStore, In: in}
	cli.PlayPoker()

	if len(playerStore.WinCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}

	got := playerStore.WinCalls[0]
	want := "Chris"

	if got != want {
		t.Errorf("didn't record correct winner, got %q, want %q", got, want)
	}
}
