package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	// %q -> wraps your values in double quotes
	// ref: https://pkg.go.dev/fmt#hdr-Printing
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
