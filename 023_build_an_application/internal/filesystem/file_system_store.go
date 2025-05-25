package filesystem

import (
	"encoding/json"
	"io"
	"log"

	"github.com/LPvdT/go-with-tests/application/common"
)

type FileSystemPlayerStore struct {
	Database io.ReadWriteSeeker
	league   common.League
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, io.SeekStart)
	league, _ := common.NewLeague(database)

	return &FileSystemPlayerStore{
		Database: database,
		league:   league,
	}
}

func (f *FileSystemPlayerStore) GetLeague() common.League {
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

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

	// Reset the reader to the start of the file
	if _, err := f.Database.Seek(0, io.SeekStart); err != nil {
		log.Fatalf("could not seek to start of database: %v", err)
	}

	// Write the updated league back to the file
	if err := json.NewEncoder(f.Database).Encode(f.league); err != nil {
		log.Fatalf("could not write to database: %v", err)
	}
}
