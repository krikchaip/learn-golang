package osexec_test

import (
	osexec "23-os-exec"
	"testing"
)

func TestGetData(t *testing.T) {
	got := osexec.GetData()
	want := "HAPPY NEW YEAR!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
