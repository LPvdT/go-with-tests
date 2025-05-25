package filesystem

import (
	"strings"
	"testing"

	"github.com/LPvdT/go-with-tests/application/common"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		store := FileSystemPlayerStore{database}

		got := store.GetLeague()

		want := []common.Player{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 33},
		}
		common.AssertLeague(t, got, want)
	})
}
