package filesystem

import (
	"io"

	"github.com/LPvdT/go-with-tests/application/common"
)

type FileSystemPlayerStore struct {
	database io.Reader
}

func (f *FileSystemPlayerStore) GetLeague() []common.Player {
	return nil
}
