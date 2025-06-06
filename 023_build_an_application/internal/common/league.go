// Package common provides common types and functions used across the application.
package common

import (
	"encoding/json"
	"fmt"
	"io"
)

// Player represents a player in the league with a name and number of wins.
type Player struct {
	Name string
	Wins int
}

// League stores a collection of players.
type League []Player

// Find tries to return a player from a league.
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

// NewLeague creates a league from JSON.
func NewLeague(rdr io.Reader) (League, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}
