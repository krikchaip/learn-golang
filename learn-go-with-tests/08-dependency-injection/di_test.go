package dependency_injection

import (
	"bytes" // ?? analogous to the "strings" package
	"testing"
)

func TestGreet(t *testing.T) {
	// ?? bytes.Buffer implements Reader and Writer interface
	buffer := &bytes.Buffer{}

	Greet(buffer, "Chris")
	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
