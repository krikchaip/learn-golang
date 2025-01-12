package assert

import "testing"

func Equal[T comparable](t *testing.T, got, want T) {
	// when t.Errorf() is called from our Equal() function,
	// the Go test runner will report the filename and line number
	// of the code **which called** our Equal() function in the output.
	t.Helper()

	if got != want {
		t.Errorf("got %#v; want %#v", got, want)
	}
}
