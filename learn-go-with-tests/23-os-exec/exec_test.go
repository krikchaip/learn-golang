package osexec_test

import (
	osexec "23-os-exec"
	"strings"
	"testing"
)

func TestGetData(t *testing.T) {
	input := strings.NewReader(`
		<payload>
			<message>Happy New Year!</message>
		</payload>
	`)

	got := osexec.GetData(input)
	want := "HAPPY NEW YEAR!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
