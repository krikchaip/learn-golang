package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	humanDate, ok := functions["humanDate"].(func(time.Time) string)

	if !ok {
		t.Fatal("humanDate template function not defined!")
		return
	}

	t.Run("17 Mar 2024 at 10:15", func(t *testing.T) {
		date := time.Date(2024, time.March, 17, 10, 15, 0, 0, time.UTC)

		got := humanDate(date)
		want := "17 Mar 2024 at 10:15"

		if got != want {
			t.Errorf("got %q; want %q", got, want)
		}
	})
}
