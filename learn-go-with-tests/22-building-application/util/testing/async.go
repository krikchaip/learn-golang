package testing

import (
	tt "testing"
	"time"
)

func Within(t tt.TB, d time.Duration, assert func()) {
	t.Helper()

	done := make(chan struct{})

	go func() {
		assert()
		done <- struct{}{}
	}()

	select {
	case <-time.After(d):
		t.Fatal("timed out after", d)
	case <-done:
	}
}
