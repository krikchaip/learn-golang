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
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error but got none.")
		}

		assertStrings(t, err.Error(), want, given)
	})
}

func assertStrings(t testing.TB, got, want, given string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, given)
	}
}
