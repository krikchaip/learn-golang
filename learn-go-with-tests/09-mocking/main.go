package main

import (
	countdown "09-mocking/lib"
	"os"
	"time"
)

// implements: countdown.Sleeper
type SecondSleeper struct {
	duration time.Duration
}

func (ss SecondSleeper) Sleep() {
	time.Sleep(ss.duration * time.Second)
}

func main() {
	writer := os.Stdout
	sleeper := SecondSleeper{1}
	countdown.Countdown(writer, sleeper)
}
