package filesystem

import (
	"testing"

	"github.com/LPvdT/go-with-tests/application/common"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := common.CreateTempFile(t, `[
			{"Name": "Chris", "Wins": 33},
			{"Name": "Cleo", "Wins": 10}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		if err != nil {
			t.Fatalf("problem creating file system player store, %v", err)
		}

		got := store.GetLeague()
		want := []common.Player{
			{Name: "Chris", Wins: 33},
			{Name: "Cleo", Wins: 10},
		}

		assertLeague(t, got, want)

		// Read the file again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := common.CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		if err != nil {
			t.Fatalf("problem creating file system player store, %v", err)
		}

		got := store.GetPlayerScore("Chris")
		want := 33

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := common.CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		if err != nil {
			t.Fatalf("problem creating file system player store, %v", err)
		}
		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := common.CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		if err != nil {
			t.Fatalf("problem creating file system player store, %v", err)
		}

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1

		assertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := common.CreateTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)

		common.AssertNoError(t, err)
	})

	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := common.CreateTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		common.AssertNoError(t, err)

		got := store.GetLeague()
		want := common.League{
			{Name: "Chris", Wins: 33},
			{Name: "Cleo", Wins: 10},
		}
		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})
}

// assertScoreEquals checks that two integer scores are equal.
// It reports a test error if they are not.
func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertLeague(t testing.TB, got, want []common.Player) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("got %d players in league, want %d", len(got), len(want))
	}

	for i, player := range got {
		if player.Name != want[i].Name || player.Wins != want[i].Wins {
			t.Errorf("got player %v, want %v", player, want[i])
		}
	}
}
