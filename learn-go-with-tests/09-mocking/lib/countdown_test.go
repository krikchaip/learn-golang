package countdown_test

import (
	countdown "09-mocking/lib"
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	// ** eventhough we switch the implementation of Countdown
	// ** to the fake one, this test would still pass.
	t.Run("don't check the timing order", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpySleeper{}

		countdown.Countdown(buffer, sleeper)

		got := buffer.String()
		want := "3\n2\n1\nGo!\n"

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

		countdown.Countdown(writer, sleeper)

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

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}

	sleeper := countdown.ConfigurableSleeper{sleepTime, spyTime.SleepFn}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
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

// ** this struct does not implement countdown.Sleeper!!
// ** pay attention to the signature of SpyTime.SleepFn
type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) SleepFn(duration time.Duration) {
	s.durationSlept = duration
}
