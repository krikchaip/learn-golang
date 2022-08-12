package racer_test

import (
	racer "11-select"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// ** [NOT RECOMMENDED] we're reaching out to real websites here !!
func TestRacer_flaky(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.dev"

	got := racer.Racer(slowURL, fastURL)
	want := fastURL

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// ** [RECOMMENDED] using a mock HTTP server instead
func TestRacer_robust(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	defer func() {
		slowServer.Close()
		fastServer.Close()
	}()

	got := racer.Racer(slowServer.URL, fastServer.URL)
	want := fastServer.URL

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
