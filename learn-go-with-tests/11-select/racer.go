package racer

import (
	"errors"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, err error) {
	winner, err = sequentRace(a, b)
	// winner, err = concurrentRace(a, b)
	return
}

// ** [NOT RECOMMENDED]
// **   - testing the speeds one after another
// **   - we measure the response times outselves
func sequentRace(a, b string) (winner string, err error) {
	durationA := measureTime(func() {
		http.Get(a)
	})

	if durationA > 10*time.Second {
		err = errors.New("timeout!")
		return
	}

	durationB := measureTime(func() {
		http.Get(b)
	})

	if durationA+durationB > 10*time.Second {
		err = errors.New("timeout!")
		return
	}

	if durationA > durationB {
		winner = b
	} else {
		winner = a
	}

	return
}

func measureTime(f func()) time.Duration {
	start := time.Now()
	f()
	return time.Since(start)
}

// ?? [RECOMMENDED]
// ??   - checking both urls at the same time
// ??   - we just only want to know which one comes back first
func concurrentRace(a, b string) (winner string) {
	select { // ** similar to Promise.race in Javascript 👍🏻
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan struct{} {
	// ** Q: Why not chan bool?
	// ** A: mem(chan struct{}) < mem(chan bool)
	ch := make(chan struct{})

	go func() {
		http.Get(url)

		// ** any <-ch expression after close(ch) will succeed without blocking,
		// ** but the result will be zero-value and ok is false
		// ** eg. result, ok := <-ch
		close(ch)
	}()

	return ch
}
