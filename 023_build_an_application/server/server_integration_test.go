package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LPvdT/go-with-tests/application/common"
	"github.com/LPvdT/go-with-tests/application/internal/filesystem"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := common.CreateTempFile(t, "")
	defer cleanDatabase()

	store := &filesystem.FileSystemPlayerStore{
		Database: database,
	}
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), common.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), common.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), common.NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, common.NewGetScoreRequest(player))

		common.AssertStatus(t, response.Code, http.StatusOK)
		common.AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, common.NewLeagueRequest())

		common.AssertStatus(t, response.Code, http.StatusOK)

		got := common.GetLeagueFromResponse(t, response.Body)
		want := []common.Player{
			{Name: "Pepper", Wins: 3},
		}
		common.AssertLeague(t, got, want)
	})
}
