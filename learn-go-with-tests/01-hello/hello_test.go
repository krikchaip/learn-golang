package hello

import "testing"

func TestHello(t *testing.T) {
	// testing.TB is an interface that holds both
	// testing.T (test helper functions) and
	// testing.B (for benchmark)
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// ?? comment the line below, make a test fail and observe the output :)
		t.Helper()

		// %q -> wraps your values in double quotes
		// ref: https://pkg.go.dev/fmt#hdr-Printing
		if got != want {
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
