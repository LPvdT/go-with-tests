package playertest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"sort"
	"testing"

	"github.com/LPvdT/go-with-tests/application/internal/common"
)

// StubPlayerStore is a fake implementation of a player store for testing.
type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   []common.Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.Scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

func (s *StubPlayerStore) GetLeague() common.League {
	return s.League
}

// AssertLeague compares two slices of Player and reports an error if they don't match.
func AssertLeague(t testing.TB, got, want []common.Player) {
	t.Helper()

	sort.Slice(got, func(i, j int) bool {
		return got[i].Name < got[j].Name
	})
	sort.Slice(want, func(i, j int) bool {
		return want[i].Name < want[j].Name
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// AssertPlayerWin checks that RecordWin was called with the correct winner.
func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()
	if len(store.WinCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
	}
	if store.WinCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.WinCalls[0], winner)
	}
}

// NewGetScoreRequest returns a GET /players/{name} request.
func NewGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

// NewPostWinRequest returns a POST /players/{name} request.
func NewPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

// NewLeagueRequest returns a GET /league request.
func NewLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

// GetLeagueFromResponse decodes the response body into a slice of Player.
func GetLeagueFromResponse(t testing.TB, body io.Reader) (league []common.Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)
	if err != nil {
		t.Fatalf("Unable to parse response from server: %v", err)
	}
	return
}
