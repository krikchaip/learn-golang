package game

import (
	"bytes"
	"slices"
	"testing"
)

func TestParseReader(t *testing.T) {
	file := bytes.NewBufferString(`
		1+1,2
		2+2,4
		10*2,20`)

	quizzes := []quiz{
		{"1+1", "2"},
		{"2+2", "4"},
		{"10*2", "20"},
	}

	t.Run("parse CSV file", func(t *testing.T) {
		var in, out bytes.Buffer
		g := New(&in, &out)

		if err := g.ParseReader(file); err != nil {
			t.Fatal(err)
		}

		if got, want := len(g.quizzes), len(quizzes); got != want {
			t.Errorf("got quizzes length of %d; want %d", got, want)
		}

		if !slices.Equal(g.quizzes, quizzes) {
			t.Errorf("got %v; want %v", g.quizzes, quizzes)
		}
	})

	t.Run("shuffle quizzes order", func(t *testing.T) {
	})
}

func TestStart(t *testing.T) {
}
