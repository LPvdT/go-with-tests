package filesystem

import (
	"encoding/json"
	"io"
	"log"

	"github.com/LPvdT/go-with-tests/application/common"
)

type FileSystemPlayerStore struct {
	database io.Reader
}

func (f *FileSystemPlayerStore) GetLeague() []common.Player {
	var league []common.Player
	if err := json.NewDecoder(f.database).Decode(&league); err != nil {
		log.Fatalf("Could not decode league from file system: %v", err)
	}
	return league
}
