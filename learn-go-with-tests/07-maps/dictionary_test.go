package maps

import "testing"

func TestSearch(t *testing.T) {
	// ?? map.key type must be conforms `comparable`
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		given := "test"
		got, _ := dictionary.Search(given)
		want := "this is just a test"

		assertStrings(t, got, want, given)
	})

	t.Run("unknown word", func(t *testing.T) {
		given := "unknown"
		_, err := dictionary.Search(given)
		want := ErrNotFound

		assertError(t, err, want)
	})
}

func assertStrings(t testing.TB, got, want, given string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, given)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("expected to get an error but got none.")
	}

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
