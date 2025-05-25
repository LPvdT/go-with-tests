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
	// Reset the reader to the start of the file
	if _, err := f.Database.Seek(0, io.SeekStart); err != nil {
		log.Fatalf("could not seek to start of database: %v", err)
	}
	league, _ := common.NewLeague(f.Database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		league = append(league, common.Player{
			Name: name,
			Wins: 1,
		})
	}

	// Reset the reader to the start of the file
	if _, err := f.Database.Seek(0, io.SeekStart); err != nil {
		log.Fatalf("could not seek to start of database: %v", err)
	}

	// Write the updated league back to the file
	if err := json.NewEncoder(f.Database).Encode(league); err != nil {
		log.Fatalf("could not write to database: %v", err)
	}
}
