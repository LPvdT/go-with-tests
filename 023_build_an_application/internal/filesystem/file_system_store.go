package filesystem

import (
	"encoding/json"
	"io"
	"log"

	"github.com/LPvdT/go-with-tests/application/common"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []common.Player {
	// Reset the reader to the start of the file
	if _, err := f.database.Seek(0, io.SeekStart); err != nil {
		log.Fatalf("could not seek to start of database: %v", err)
	}
	league, _ := common.NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int

	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}

	return wins
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()

	for i, player := range league {
		if player.Name == name {
			league[i].Wins++
		}
	}

	// Reset the reader to the start of the file
	if _, err := f.database.Seek(0, io.SeekStart); err != nil {
		log.Fatalf("could not seek to start of database: %v", err)
	}

	// Write the updated league back to the file
	if err := json.NewEncoder(f.database).Encode(league); err != nil {
		log.Fatalf("could not write to database: %v", err)
	}
}
