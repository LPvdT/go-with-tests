package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LPvdT/go-with-tests/application/common"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []common.Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() common.League {
	return s.league
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := common.NewGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		common.AssertStatus(t, response.Code, http.StatusOK)
		common.AssertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := common.NewGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		common.AssertStatus(t, response.Code, http.StatusOK)
		common.AssertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := common.NewGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		common.AssertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"

		request := common.NewPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		common.AssertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []common.Player{
			{Name: "Cleo", Wins: 32},
			{Name: "Chris", Wins: 20},
			{Name: "Tiest", Wins: 14},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		request := common.NewLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := common.GetLeagueFromResponse(t, response.Body)

		common.AssertStatus(t, response.Code, http.StatusOK)
		common.AssertLeague(t, got, wantedLeague)
		common.AssertContentType(t, response, jsonContentType)
	})
}
