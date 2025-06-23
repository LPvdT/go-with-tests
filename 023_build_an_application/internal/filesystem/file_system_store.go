package filesystem

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/LPvdT/go-with-tests/application/common"
)

// FileSystemPlayerStore stores players in the filesystem.
type FileSystemPlayerStore struct {
	Database *json.Encoder
	league   common.League
}

// NewFileSystemPlayerStore creates a FileSystemPlayerStore initialising the store if needed.
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initialisePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := common.NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		Database: json.NewEncoder(&common.Tape{File: file}),
		league:   league,
	}, nil
}

// FileSystemPlayerStoreFromFile creates a PlayerStore from the contents of a JSON file found at path.
func FileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0o666)
	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s %v", path, err)
	}

	closeFunc := func() {
		_ = db.Close()
	}

	store, err := NewFileSystemPlayerStore(db)
	if err != nil {
		return nil, nil, fmt.Errorf("problem creating file system player store, %v ", err)
	}

	return store, closeFunc, nil
}

func initialisePlayerDBFile(file *os.File) error {
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		_, err = file.Write([]byte("[]"))
		if err != nil {
			return err
		}
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetLeague returns the Scores of all the players.
func (f *FileSystemPlayerStore) GetLeague() common.League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

// GetPlayerScore retrieves a player's score.
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

// RecordWin will store a win for a player, incrementing wins if already known.
func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, common.Player{Name: name, Wins: 1})
	}

	_ = f.Database.Encode(f.league)
}
