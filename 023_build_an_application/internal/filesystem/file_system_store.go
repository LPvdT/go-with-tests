package filesystem

import (
	"io"
	"log"

	"github.com/LPvdT/go-with-tests/application/common"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
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
	return 0
}
