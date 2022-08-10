package main

import (
	countdown "09-mocking/lib"
	"os"
)

func main() {
	countdown.Countdown(os.Stdout)
}
