package sync_test

import (
	lib "13-sync"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times, leaves it at 3", func(t *testing.T) {
		// var counter lib.Counter
		counter := lib.NewCounter()

		counter.Inc()
		counter.Inc()
		counter.Inc()

		// assertCounter__danger(t, counter, 3)
		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		// var counter lib.Counter
		counter := lib.NewCounter()

		var routines sync.WaitGroup
		count := 1000

		// ** set the number of goroutines to wait for.
		routines.Add(count)

		for i := 0; i < count; i++ {
			go func() {
				// ** decrement WG counter by one
				defer routines.Done()

				counter.Inc()
			}()
		}

		// ** will wait(block) until all goroutines have finished
		routines.Wait()

		// assertCounter__danger(t, counter, count)
		assertCounter(t, counter, count)
	})
}

// ?? run `go vet [package.go]` to see the warning
// ** "A MUTEX MUST NOT BE COPIED BY VALUE AFTER FIRST USE"
func assertCounter__danger(t testing.TB, counter lib.Counter, want int) {
	t.Helper()
	if counter.Value() != want {
		t.Errorf("got %d, want %d", counter.Value(), want)
	}
}

func assertCounter(t testing.TB, counter *lib.Counter, want int) {
	t.Helper()
	if counter.Value() != want {
		t.Errorf("got %d, want %d", counter.Value(), want)
	}
}
