package countdown

import (
	"bytes"
	"testing"
)

// implements: countdown.Sleeper
type SpySleeper struct {
	Calls int
}

func (ss *SpySleeper) Sleep() {
	ss.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := &SpySleeper{}

	Countdown(buffer, sleeper)

	got := buffer.String()
	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}

	if sleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", sleeper.Calls)
	}
}
