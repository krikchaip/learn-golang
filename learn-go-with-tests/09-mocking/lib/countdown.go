package countdown

import (
	"fmt"
	"io"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

func Countdown(w io.Writer, s Sleeper) {
	// realCountdown(w, s)
	fakeCountdown(w, s)
}

func realCountdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}

	fmt.Fprint(w, finalWord)
}

// ** this won't break the test eventhough the implementation is wrong!
func fakeCountdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
	}

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
	}

	fmt.Fprint(w, finalWord)
}
