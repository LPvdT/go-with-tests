package testutils

import "testing"

// AssertScoreEquals checks that two integer scores are equal.
// It reports a test error if they are not.
func AssertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
