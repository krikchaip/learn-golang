package main

import (
	"krikchaip/quiz-game/internal"
	"krikchaip/quiz-game/internal/options"
	"log"
	"os"
)

func main() {
	// parse command line options for the quiz game
	options.Parse()

	// initiate quiz game instance
	g := game.New(
		os.Stdin, os.Stdout,
		game.WithLimit(options.Values.Limit),
		game.WithShuffle(options.Values.Shuffle),
	)

	problems := readProblemFile()
	defer problems.Close()

	if err := g.ParseReader(problems); err != nil {
		log.Fatal(err)
	}

	// game start
	if err := g.Start(); err != nil {
		log.Fatal(err)
	}
}

func readProblemFile() *os.File {
	problems, err := os.Open(options.Values.CSV)
	if err != nil {
		log.Fatal(err)
	}

	return problems
}
