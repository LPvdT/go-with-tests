package store

import "github.com/LPvdT/go-with-tests/application/common"

// InMemoryPlayerStore collects data about players in memory.
type InMemoryPlayerStore struct {
	store map[string]int
}

// NewInMemoryPlayerStore initialises an empty player store.
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

// GetLeague returns a collection of Players.
func (i *InMemoryPlayerStore) GetLeague() common.League {
	var league []common.Player
	for name, wins := range i.store {
		league = append(league, common.Player{Name: name, Wins: wins})
	}
	return league
}
