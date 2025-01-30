package game

import (
	"encoding/csv"
	"errors"
	"io"
)

type quiz struct{ question, answer string }

type game struct {
	quizzes []quiz
	scores  uint
}

func New() *game {
	return &game{}
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
