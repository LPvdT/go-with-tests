package main

import (
	"strings"
	"testing"

	"github.com/LPvdT/go-with-tests/application/internal/cli"
	"github.com/LPvdT/go-with-tests/application/internal/common"
)

func TestCLI(t *testing.T) {
	in := strings.NewReader(("Chris wins\n"))
	playerStore := &common.StubPlayerStore{}

	cli := &cli.CLI{PlayerStore: playerStore, In: in}
	cli.PlayPoker()

	common.AssertPlayerWin(t, playerStore, "Chris")
}
