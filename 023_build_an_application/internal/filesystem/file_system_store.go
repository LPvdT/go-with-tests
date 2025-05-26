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

	err := f.Database.Encode(f.league)
	if err != nil {
		log.Fatalf("could not encode league to file: %v", err)
	}
}
