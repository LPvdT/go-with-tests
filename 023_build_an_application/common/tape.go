package common

import (
	"io"
	"os"
)

type Tape struct {
	File *os.File
}

func (t *Tape) Write(p []byte) (n int, err error) {
	if err := t.File.Truncate(0); err != nil {
		return 0, err
	}

	if _, err := t.File.Seek(0, io.SeekStart); err != nil {
		return 0, err
	}

	return t.File.Write(p)
}
