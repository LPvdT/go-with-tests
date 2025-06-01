package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LPvdT/go-with-tests/application/internal/common"
	"github.com/LPvdT/go-with-tests/application/internal/filesystem"
	"github.com/LPvdT/go-with-tests/application/playertest"
	"github.com/LPvdT/go-with-tests/application/testutils"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := testutils.CreateTempFile(t, `[]`)
	defer cleanDatabase()

	store := &filesystem.FileSystemPlayerStore{
		Database: json.NewEncoder(&common.Tape{File: database}),
	}

	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), playertest.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), playertest.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), playertest.NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, playertest.NewGetScoreRequest(player))

		testutils.AssertStatus(t, response.Code, http.StatusOK)
		testutils.AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, playertest.NewLeagueRequest())
		testutils.AssertStatus(t, response.Code, http.StatusOK)

		got := playertest.GetLeagueFromResponse(t, response.Body)
		want := []common.Player{
			{Name: "Pepper", Wins: 3},
		}
		playertest.AssertLeague(t, got, want)
	})
}
