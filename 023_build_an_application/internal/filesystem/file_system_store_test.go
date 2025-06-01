package filesystem

import (
	"testing"

	"github.com/LPvdT/go-with-tests/application/internal/common"
	"github.com/LPvdT/go-with-tests/application/playertest"
	"github.com/LPvdT/go-with-tests/application/testutils"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := testutils.CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		if err != nil {
			t.Fatalf("problem creating file system player store, %v", err)
		}

		got := store.GetLeague()
		want := []common.Player{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 33},
		}

		playertest.AssertLeague(t, got, want)

		// Read the file again
		got = store.GetLeague()
		playertest.AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := testutils.CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		if err != nil {
			t.Fatalf("problem creating file system player store, %v", err)
		}

		got := store.GetPlayerScore("Chris")
		want := 33

		testutils.AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := testutils.CreateTempFile(t, `[
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

		testutils.AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := testutils.CreateTempFile(t, `[
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

		testutils.AssertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := testutils.CreateTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)

		testutils.AssertNoError(t, err)
	})

	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := testutils.CreateTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		testutils.AssertNoError(t, err)

		got := store.GetLeague()
		want := common.League{
			{Name: "Chris", Wins: 33},
			{Name: "Cleo", Wins: 10},
		}
		playertest.AssertLeague(t, got, want)

		got = store.GetLeague()
		playertest.AssertLeague(t, got, want)
	})
}
