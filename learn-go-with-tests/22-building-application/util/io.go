package util

import (
	"io"
	"os"
)

// Encapsulates the "read/rewrite from the beginning" logic
//
// implements: io.ReadWriter
type Tape struct {
	file *os.File
}

func NewTape(source *os.File) *Tape {
	return &Tape{source}
}

func (t *Tape) Read(p []byte) (n int, err error) {
	t.file.Seek(0, io.SeekStart)
	return t.file.Read(p)
}

func (t *Tape) Write(p []byte) (n int, err error) {
	// It is possible that the file cursor might reached the EOF already.
	// We seek from the beginning to make sure it's still fresh.
	t.file.Seek(0, io.SeekStart)

	// delete all contents before rewriting
	t.file.Truncate(0)

	return t.file.Write(p)
}
