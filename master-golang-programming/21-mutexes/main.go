package main

import (
	"master-golang-programming/url-downloader/lib"
	"sync"
	"time"
)

const Groups int = 100

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	// ?? 2 goroutines are running simulteneously in each loop
	wg.Add(Groups * 2)

	// ?? shared variable (state)
	var n int = 0

	for i := 0; i < Groups; i++ {
		go func() {
			time.Sleep(100 * time.Millisecond)
			atomic(&m, func() { n++; wg.Done() })
		}()

		go func() {
			time.Sleep(200 * time.Millisecond)
			atomic(&m, func() { n--; wg.Done() })
		}()
	}

	// ?? waiting for all goroutines to finish
	lib.MeasureTime(func() { wg.Wait() }) // should be around 200ms

	println(n) // always = 0
}

func atomic(mut *sync.Mutex, fn func()) {
	mut.Lock()
	fn()
	mut.Unlock()
}
