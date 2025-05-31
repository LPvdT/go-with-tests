package cli

import "github.com/LPvdT/go-with-tests/application/internal/server"

type CLI struct {
	PlayerStore server.PlayerStore
}

func (cli *CLI) PlayPoker() {
	cli.PlayerStore.RecordWin("Cleo")
}
