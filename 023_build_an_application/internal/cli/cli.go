package cli

import (
	"io"

	"github.com/LPvdT/go-with-tests/application/internal/server"
)

type CLI struct {
	PlayerStore server.PlayerStore
	In          io.Reader
}

func (cli *CLI) PlayPoker() {
	cli.PlayerStore.RecordWin("Cleo")
}
