package lib

import (
	"fmt"
	"sync"
	"time"
)

func Public(path string) string {
	if rune(path[0]) != '/' {
		path = "/" + path
	}

	return "./public" + path
}

func Concurrently[T, U any](tasks []T, fn func(T) U) {
	var wg sync.WaitGroup

	wg.Add(len(tasks))

	for _, t := range tasks {
		go func(t T) {
			fn(t)
			wg.Done()
		}(t)
	}

	wg.Wait()
}

func MeasureTime(f func()) {
	start := time.Now()
	f()
	used := time.Since(start)
	fmt.Println("time used:", used.Seconds(), "seconds")
}
