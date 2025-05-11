package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LPvdT/go-with-tests/server"
)

func TestGETPlayers(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	server.PlayerServer(response, request)

	t.Run("GET / should return 200 OK", func(t *testing.T) {
		got := response.Body.String()
		want := "200 OK"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
