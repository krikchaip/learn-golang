package main

import (
	"os"
	"time"

	cd "09-mocking/lib"
)

func main() {
	writer := os.Stdout

	var sleeper cd.Sleeper

	sleeper = &cd.SecondSleeper{Duration: 1}
	cd.Countdown(writer, sleeper)

	sleeper = &cd.ConfigurableSleeper{
		Duration: 500 * time.Millisecond,
		SleepFn:  time.Sleep,
	}
	cd.Countdown(writer, sleeper)
}
