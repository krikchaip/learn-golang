package game

import (
	"bytes"
	"slices"
	"testing"
)

func TestParseReader(t *testing.T) {
	t.Run("parse CSV file", func(t *testing.T) {
		var in, out bytes.Buffer
		g := New(&in, &out)

		file, quizzes := generateTestSample()

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
		var in, out bytes.Buffer
		g := New(&in, &out, WithShuffle(true))

		file, quizzes := generateTestSample()

		if err := g.ParseReader(file); err != nil {
			t.Fatal(err)
		}

		if got, want := len(g.quizzes), len(quizzes); got != want {
			t.Errorf("got quizzes length of %d; want %d", got, want)
		}

		if slices.Equal(g.quizzes, quizzes) {
			t.Errorf("the results order are expected to be shuffled")
		}
	})
}

func TestStart(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		scores uint
	}{
		{name: "finish the game successfully", input: "2\n4\n20\n", scores: 3},
		{name: "got only 2 scores", input: "2\n\n20\n", scores: 2},
		{name: "skip answers", input: "\n\n\n", scores: 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var in, out bytes.Buffer
			g := New(&in, &out, WithLimit(10))

			file, _ := generateTestSample()

			if err := g.ParseReader(file); err != nil {
				t.Fatal(err)
			}

			// start the quiz game asynchronously
			finished := make(chan error)
			go func() {
				finished <- g.Start()
				close(finished)
			}()

			// write quiz answers
			in.WriteString(tc.input)

			if err := <-finished; err != nil {
				t.Fatal(err)
			}

			if got, want := g.scores, tc.scores; got != want {
				t.Errorf("got scores %d; want %d", got, want)
			}
		})
	}
}

func generateTestSample() (file *bytes.Buffer, quizzes []quiz) {
	file = bytes.NewBufferString(`
		1+1,2
		2+2,4
		10*2,20`)

	quizzes = []quiz{
		{"1+1", "2"},
		{"2+2", "4"},
		{"10*2", "20"},
	}

	return
}
