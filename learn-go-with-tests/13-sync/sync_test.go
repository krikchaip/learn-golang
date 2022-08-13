package sync_test

import (
	lib "13-sync"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times, leaves it at 3", func(t *testing.T) {
		var counter lib.Counter

		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value() != 3 {
			t.Errorf("got %d, want %d", counter.Value(), 3)
		}
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		var counter lib.Counter
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

		if counter.Value() != count {
			t.Errorf("got %d, want %d", counter.Value(), count)
		}
	})
}
