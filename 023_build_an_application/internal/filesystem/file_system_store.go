// Package filesystem provides a file system implementation of the PlayerStore interface.
package filesystem

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/LPvdT/go-with-tests/application/internal/common"
)

type FileSystemPlayerStore struct {
	Database *json.Encoder
	league   common.League
}

// NewFileSystemPlayerStore creates a new PlayerStore that stores data in a file.
//
// The file is expected to be a JSON file containing the current league.
// The file is read when the store is created and written to when the store is updated.
func NewFileSystemPlayerStore(file *os.File) *FileSystemPlayerStore {
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatalf("could not seek to start of file: %v", err)
	}

	league, _ := common.NewLeague(file)

	return &FileSystemPlayerStore{
		Database: json.NewEncoder(&common.Tape{File: file}),
		league:   league,
	}
}

// GetLeague retrieves the current league state from the file.
func (f *FileSystemPlayerStore) GetLeague() common.League {
	return f.league
}

// GetPlayerScore returns the number of wins for a given player.
// If the player is not found in the league, it returns 0.
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

// RecordWin increments the win count for a player with the given name.
//
// If the player is not found, it adds a new player with a win count of 1.
// After the win count is updated, it writes the updated league to the file.
func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, common.Player{
			Name: name,
			Wins: 1,
		})
	}

	err := f.Database.Encode(f.league)
	if err != nil {
		log.Fatalf("could not encode league to file: %v", err)
	}
}
