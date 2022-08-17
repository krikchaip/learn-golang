package clockface_test

import (
	clockface "16-maths"
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {
	t.Run("at midnight", func(t *testing.T) {
		tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

		got := clockface.SecondHand(tm)
		want := clockface.Point{X: 150, Y: 60}

		if got != want {
			t.Errorf("Got %v, wanted %v", got, want)
		}
	})

	t.Run("at 30 seconds", func(t *testing.T) {
		tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

		got := clockface.SecondHand(tm)
		want := clockface.Point{X: 150, Y: 240}

		if got != want {
			t.Errorf("Got %v, wanted %v", got, want)
		}
	})
}
