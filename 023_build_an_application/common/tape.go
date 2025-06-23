package common

import (
	"io"
	"os"
)

type Tape struct {
	File *os.File
}

// Write overwrites the contents of the file with the given []byte. If
// there is an error truncating or seeking the file, it is returned.
func (t *Tape) Write(p []byte) (n int, err error) {
	if err := t.File.Truncate(0); err != nil {
		return 0, err
	}

	if _, err := t.File.Seek(0, io.SeekStart); err != nil {
		return 0, err
	}

	return t.File.Write(p)
}
