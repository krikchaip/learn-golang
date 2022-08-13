package main

import (
	"fmt"
	"master-golang-programming/url-downloader/lib"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://www.facebook.com",
		"https://www.google.com",
		"https://www.medium.com",
	}

	measureTime(func() {
		// for _, url := range urls {
		// 	lib.DownloadURL(url)
		// }

		concurrently(urls, lib.DownloadURL)
	})
}

func concurrently[T, U any](tasks []T, fn func(T) U) {
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

func measureTime(f func()) {
	start := time.Now()
	f()
	used := time.Since(start)
	fmt.Println("time used:", used.Seconds(), "seconds")
}
