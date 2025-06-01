package common

import (
	"io"
	"testing"

	"github.com/LPvdT/go-with-tests/application/testutils"
)

func TestTape(t *testing.T) {
	file, clean := testutils.CreateTempFile(t, "12345")
	defer clean()

	tape := &Tape{file}

	_, err := tape.Write([]byte("abc"))
	if err != nil {
		t.Errorf("got %v want nil", err)
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		t.Fatalf("could not seek to start of file: %v", err)
	}

	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
