package countdown

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	// ** eventhough we switch the implementation of Countdown
	// ** to the fake one, this test would still pass.
	t.Run("don't check the timing order", func(t *testing.T) {
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
	})

	t.Run("sleep before every print", func(t *testing.T) {
		writer := &SpyCountdownOperations{}
		sleeper := writer

		Countdown(writer, sleeper)

		got := writer.Calls // ??  or sleeper.Calls
		want := []string{
			Write, // 3
			Sleep,
			Write, // 2
			Sleep,
			Write, // 1
			Sleep,
			Write, // Go!
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted calls %v but got %v", want, got)
		}
	})
}

// implements: countdown.Sleeper
type SpySleeper struct {
	Calls int
}

func (ss *SpySleeper) Sleep() {
	ss.Calls++
}

// implements: io.Writer, countdown.Sleeper
type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, Sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, Write)
	return
}

const (
	Write = "write"
	Sleep = "sleep"
)
