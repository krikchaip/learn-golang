package game

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"
)

type quiz struct{ question, answer string }

type game struct {
	input   *bufio.Scanner
	output  io.Writer
	timeout time.Duration

	quizzes []quiz
	scores  uint
}

func New(r io.Reader, w io.Writer, limit uint) *game {
	input := bufio.NewScanner(r)
	timeout := time.Duration(limit) * time.Second

	return &game{input: input, output: w, timeout: timeout}
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

		g.quizzes = append(g.quizzes, quiz{record[0], record[1]})
	}

	return nil
}

func (g *game) Start() error {
	for i, q := range g.quizzes {
		fmt.Fprintf(g.output, "Problem #%d: %s = ", i+1, q.question)

		line, err := g.readLine()
		if err != nil {
			return err
		}

		if line == q.answer {
			g.scores++
		}
	}

	fmt.Fprintf(g.output, "You scored %d out of %d.", g.scores, len(g.quizzes))

	return nil
}

func (g *game) readLine() (string, error) {
	if ok, err := g.input.Scan(), g.input.Err(); !ok && err != nil {
		return "", err
	}

	return strings.TrimSpace(g.input.Text()), nil
}
