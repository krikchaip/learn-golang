package game

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"
)

type quiz struct{ question, answer string }

type game struct {
	input   *bufio.Scanner
	output  io.Writer
	timeout time.Duration
	shuffle bool

	quizzes []quiz
	scores  uint
}

func New(r io.Reader, w io.Writer, options ...gameOption) *game {
	input := bufio.NewScanner(r)

	g := &game{input: input, output: w}

	// apply options
	for _, o := range options {
		o(g)
	}

	return g
}

type gameOption = func(*game)

func WithLimit(limit uint) gameOption {
	return func(g *game) {
		g.timeout = time.Duration(limit) * time.Second
	}
}

func WithShuffle(enable bool) gameOption {
	return func(g *game) {
		g.shuffle = enable
	}
}

func (g *game) ParseReader(r io.Reader) error {
	content := csv.NewReader(r)

	for {
		record, err := content.Read()

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return err
		}

		g.quizzes = append(g.quizzes, quiz{
			strings.TrimSpace(record[0]),
			strings.TrimSpace(record[1]),
		})
	}

	if g.shuffle {
		rand.Shuffle(len(g.quizzes), func(i, j int) {
			g.quizzes[i], g.quizzes[j] = g.quizzes[j], g.quizzes[i]
		})
	}

	return nil
}

func (g *game) Start() (err error) {
	select {
	case <-time.After(g.timeout):
		fmt.Fprintln(g.output)
	case err = <-g.play():
	}

	if err != nil {
		return err
	}

	fmt.Fprintf(g.output, "You scored %d out of %d.", g.scores, len(g.quizzes))

	return nil
}

func (g *game) play() <-chan error {
	sig := make(chan error)

	go func() {
		for i, q := range g.quizzes {
			fmt.Fprintf(g.output, "Problem #%d: %s = ", i+1, q.question)

			line, err := g.readLine()
			if err != nil {
				sig <- err
				break
			}

			if line == q.answer {
				g.scores++
			}
		}

		close(sig)
	}()

	return sig
}

func (g *game) readLine() (string, error) {
	if ok, err := g.input.Scan(), g.input.Err(); !ok && err != nil {
		return "", err
	}

	return strings.TrimSpace(g.input.Text()), nil
}
