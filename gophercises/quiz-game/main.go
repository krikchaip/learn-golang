package main

import (
	"fmt"
	"krikchaip/quiz-game/internal/options"
)

func main() {
	// parse command line options for the quiz game
	options.Parse()

	fmt.Printf("%#v", options.Values)
}
