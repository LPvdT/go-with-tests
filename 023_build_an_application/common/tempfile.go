package common

import (
	"os"
	"testing"
)

// CreateTempFile creates a temporary file with the specified initial data and
// returns the file pointer along with a cleanup function.
func CreateTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	if _, err := tmpfile.WriteString(initialData); err != nil {
		t.Fatalf("could not write to temp file %v", err)
	}

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

// AssertNoError checks that the error is nil. If not, it reports a fatal error.
func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("got an error but didn't want one: %v", err)
	}
}
