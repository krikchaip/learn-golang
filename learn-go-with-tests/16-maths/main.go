package main

import (
	"os"
	"time"

	clockface "16-maths/lib"
)

func main() {
	clock := clockface.ClockSVG(time.Now())
	os.WriteFile("assets/clock.svg", []byte(clock), 0666)
}
