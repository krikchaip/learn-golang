package main

import (
	countdown "09-mocking/lib"
	"os"
)

func main() {
	writer := os.Stdout
	sleeper := countdown.SecondSleeper{Duration: 1}
	countdown.Countdown(writer, sleeper)
}
