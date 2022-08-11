package racer_test

import (
	racer "11-select"
	"testing"
)

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.dev"

	got := racer.Racer(slowURL, fastURL)
	want := fastURL

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
