package main

import (
	"krikchaip/snippetbox/internal/assert"
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	humanDate, ok := functions["humanDate"].(func(time.Time) string)

	if !ok {
		t.Fatal("humanDate template function not defined!")
		return
	}

	cases := []struct {
		name  string
		input time.Time
		want  string
	}{
		{
			name:  "UTC",
			input: time.Date(2024, time.March, 17, 10, 15, 0, 0, time.UTC),
			want:  "17 Mar 2024 at 10:15",
		},
		{
			name:  "Empty",
			input: time.Time{},
			want:  "",
		},
		{
			name:  "CET Timezone",
			input: time.Date(2024, time.March, 17, 11, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want:  "17 Mar 2024 at 10:15",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := humanDate(c.input)
			want := c.want

			assert.Equal(t, got, want)
		})
	}
}
