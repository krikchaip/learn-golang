package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	// winner, err = SequentRace(a, b, tenSecondTimeout)
	winner, err = ConcurrentRace(a, b, tenSecondTimeout)
	return
}

// ** [NOT RECOMMENDED]
// **   - testing the speeds one after another
// **   - we measure the response times outselves
func SequentRace(a, b string, timeout time.Duration) (winner string, err error) {
	durationA := measureTime(func() {
		http.Get(a)
	})

	if durationA > timeout {
		err = fmt.Errorf("timed out waiting for %s", a)
		return
	}

	durationB := measureTime(func() {
		http.Get(b)
	})

	if durationA+durationB > timeout {
		err = fmt.Errorf("timed out waiting for %s and %s", a, b)
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
func ConcurrentRace(a, b string, timeout time.Duration) (winner string, err error) {
	select { // ** similar to Promise.race in Javascript 👍🏻
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout): // ** both a and b are racing against a timeout
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
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
