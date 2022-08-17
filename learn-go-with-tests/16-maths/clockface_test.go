package clockface_test

import (
	clockface "16-maths"
	"testing"
	"time"
)

func TestSecondHand(t *testing.T) {
	// time -> unit vector
	cases := []struct {
		time  time.Time
		point clockface.Point
	}{
		{clock(0, 0, 0), clockface.Point{0, -1}},
		{clock(0, 0, 30), clockface.Point{0, 1}},
		{clock(0, 0, 45), clockface.Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clockface.SecondHand(c.time)
			want := c.point.ShiftLength(clockface.SecondHandLength)

			if !got.RoughlyEqual(want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func clock(hour, minute, second int) time.Time {
	return time.Date(1337, time.January, 1, hour, minute, second, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
