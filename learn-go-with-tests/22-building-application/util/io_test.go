package util_test

import (
	"io"
	"testing"

	"22-building-application/util"
	testutil "22-building-application/util/testing"
)

// What happen if the new data is smaller than what it was before 🤔
// eg. "12345" -> "abc"
func TestTape_Write(t *testing.T) {
	file, cleanup := testutil.CreateTempFile(t, "12345")
	t.Cleanup(cleanup)

	tape := util.NewTape(file)

	// should override the content of the file
	tape.Write([]byte("abc"))

	// read the content again
	file.Seek(0, io.SeekStart)
	content, _ := io.ReadAll(file)

	got := string(content)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
