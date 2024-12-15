// to test this file, execute `go test` or `go test 01-hello`
package main

import "testing"

func TestHello(t *testing.T) {
	// testing.TB is an interface that holds both
	// testing.T (test helper functions) and
	// testing.B (for benchmark)
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// to tell the test suite that this function is a helper
		// when it fails, the line number reported will be in our function call
		// rather than inside our test helper
		t.Helper()

		// %q -> wraps your values in double quotes
		// ref: https://pkg.go.dev/fmt#hdr-Printing
		if got != want {
			// prints out an error message and fail the test
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
