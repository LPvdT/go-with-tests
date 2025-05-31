package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"sort"
	"testing"
)

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.Scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.League
}

// NewGetScoreRequest returns a new http.Request for a GET /players/{name} request.
func NewGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

// AssertResponseBody checks if the actual response body matches the expected value.
// If they do not match, it reports an error with the details.
func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

// AssertStatus checks whether the actual HTTP status code matches the expected
// status code. If they do not match, it reports an error with the details.
func AssertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

// NewPostWinRequest returns a new http.Request for a POST /players/{name} request.
func NewPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

// AssertLeague compares two slices of Player and reports an error if they are not deeply equal.
// It is useful for asserting that the actual league data matches the expected data in tests.
func AssertLeague(t testing.TB, got, want []Player) {
	t.Helper()

	// The league slice (which is written by hand in the test) needs
	// to be sorted by name.
	sort.Slice(got, func(i, j int) bool {
		return got[i].Name < got[j].Name
	})

	// The expected league slice (which is written by hand in the test) needs
	// to be sorted by name too.
	sort.Slice(want, func(i, j int) bool {
		return want[i].Name < want[j].Name
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// NewLeagueRequest returns a new http.Request for a GET /league request.
func NewLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

// GetLeagueFromResponse decodes the response body into a slice of Player.
//
// It reports a fatal error if the response cannot be parsed.
// This function is a test helper, so it should be used within a test context.
func GetLeagueFromResponse(t testing.TB, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}
	return
}

// AssertContentType checks that the Content-Type header of the response is set to the expected value.
//
// It reports a test error if the header is not set to the expected value.
func AssertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}

// AssertScoreEquals checks that the got and want scores are equal.
//
// It reports a test error if the scores are not equal.
func AssertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

// CreateTempFile creates a temporary file with the specified initial data and
// returns the file pointer along with a cleanup function.
//
// The function is a test helper, and it will report a fatal error if it cannot
// create the file or write the initial data. The cleanup function should be
// called to close and remove the temporary file after use.
func CreateTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	if _, err := tmpfile.WriteString(initialData); err != nil {
		t.Fatalf("could not write to temp file %v", err)
	}

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

// AssertNoError checks that the error is nil. If not, it reports a fatal error
// with the details. This function is a test helper, so it should be used within
// a test context.
func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("got an error but didn't want one: %v", err)
	}
}

// AssertPlayerWin checks that the RecordWin function was called with the correct winner.
//
// It reports a fatal error if the function was not called with the correct winner.
// This function is a test helper, so it should be used within a test context.
func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()
	if len(store.WinCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
	}
	if store.WinCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.WinCalls[0], winner)
	}
}
