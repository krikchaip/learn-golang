package maps

import "testing"

func TestSearch(t *testing.T) {
	// ?? map.key type must be conforms `comparable`
	dictionary := map[string]string{"test": "this is just a test"}

	given := "test"
	got := Search(dictionary, given)
	want := "this is just a test"

	assertStrings(t, got, want, given)

}

func assertStrings(t testing.TB, got, want, given string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, given)
	}
}
