package common

import (
	"io"
	"log"
)

type Tape struct {
	File io.ReadWriteSeeker
}

func (t *Tape) Write(p []byte) (n int, err error) {
	if _, err := t.File.Seek(0, io.SeekEnd); err != nil {
		log.Fatalf("could not seek to end of file: %v", err)
	}
	return t.File.Write(p)
}
