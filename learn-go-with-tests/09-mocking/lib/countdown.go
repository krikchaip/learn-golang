package countdown

import (
	"fmt"
	"io"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

// ?? comment on one of the implementations below and see the test output
func Countdown(w io.Writer, s Sleeper) {
	realCountdown(w, s)
	// fakeCountdown(w, s)
}

func realCountdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}

	fmt.Fprintln(w, finalWord)
}

// ** this won't break the test eventhough the implementation is wrong!
func fakeCountdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
	}

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
	}

	fmt.Fprintln(w, finalWord)
}
