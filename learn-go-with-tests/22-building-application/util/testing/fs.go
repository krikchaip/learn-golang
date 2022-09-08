package testing

import (
	"os"
	tt "testing"
)

func CreateTempFile(t tt.TB, content string) (
	file *os.File,
	removeFile func(),
) {
	t.Helper()

	tmp, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmp.Write([]byte(content))

	file = tmp
	removeFile = func() {
		tmp.Close()
		os.Remove(tmp.Name())
	}

	return
}
