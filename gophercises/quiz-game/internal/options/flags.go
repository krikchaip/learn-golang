package options

import "flag"

func defineFlags() {
	flag.StringVar(
		&Values.CSV,
		"csv",
		"problems.csv",
		"a csv file in the format of 'question,answer'",
	)

	flag.UintVar(
		&Values.Limit,
		"limit",
		30,
		"the time limit for the quiz in seconds",
	)

	flag.BoolVar(
		&Values.Shuffle,
		"shuffle",
		false,
		"shuffle the quiz order",
	)
}

func Parse() {
	defineFlags()
	flag.Parse()
}
