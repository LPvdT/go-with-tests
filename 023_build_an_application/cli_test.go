package main

import (
	"testing"

	"github.com/LPvdT/go-with-tests/application/internal/cli"
	"github.com/LPvdT/go-with-tests/application/internal/common"
)

func TestCLI(t *testing.T) {
	playerStore := &common.StubPlayerStore{}
	cli := &cli.CLI{PlayerStore: playerStore}
	cli.PlayPoker()

	if len(playerStore.WinCalls) != 1 {
		t.Fatal("expected a call to Win() to be made, but it wasn't")
	}
}
