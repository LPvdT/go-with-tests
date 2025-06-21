// Package server provides a HTTP interface for player information.
package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/LPvdT/go-with-tests/application/internal/common"
)

const jsonContentType = "application/json"

// PlayerStore stores score information about players.
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() common.League
}

// PlayerServer is a HTTP interface for player information.
type PlayerServer struct {
	store PlayerStore
	http.Handler
}

// NewPlayerServer creates a PlayerServer with routing configured.
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router

	return p
}

// leagueHandler responds to GET requests to "/league" with the current league.
func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)

	if err := json.NewEncoder(w).Encode(p.store.GetLeague()); err != nil {
		fmt.Println(err)
	}
}

// playersHandler handles requests to "/players/{name}".
// It processes wins for players or retrieves their scores.
func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

// showScore writes the score of the specified player to the response writer.
// If the player's score is not found, it responds with a 404 status code.
func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprint(w, score)
}

// processWin records a win for a player and returns a 202 Accepted status code
func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
