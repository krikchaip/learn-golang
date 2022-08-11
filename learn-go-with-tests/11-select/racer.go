package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	durationA := measureTime(func() {
		http.Get(a)
	})

	durationB := measureTime(func() {
		http.Get(b)
	})

	if durationA > durationB {
		return b
	} else {
		return a
	}
}

func measureTime(f func()) time.Duration {
	start := time.Now()
	f()
	return time.Since(start)
}
