package main

import (
	countdown "09-mocking/lib"
	"os"
	"time"
)

func main() {
	writer := os.Stdout

	var sleeper countdown.Sleeper

	sleeper = countdown.SecondSleeper{Duration: 1}
	countdown.Countdown(writer, sleeper)

	sleeper = &countdown.ConfigurableSleeper{
		Duration: 500 * time.Millisecond,
		SleepFn:  time.Sleep,
	}
	countdown.Countdown(writer, sleeper)
}
