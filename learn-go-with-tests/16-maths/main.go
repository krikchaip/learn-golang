package main

import (
	"os"
	"time"

	clockface "16-maths/lib"
)

// go run 16-maths/main.go
func main() {
	clock := clockface.ClockSVG(time.Now())
	os.WriteFile("16-maths/assets/clock.svg", []byte(clock), 0666)
}
