package main

import (
	"master-golang-programming/url-downloader/lib"
)

func main() {
	urls := []string{
		"https://www.facebook.com",
		"https://www.google.com",
		"https://www.medium.com",
	}

	lib.MeasureTime(func() {
		// ?? sequentially - basic
		// for _, url := range urls {
		// 	lib.DownloadURL(url)
		// }

		// ?? concurrently - using WaitGroup and Mutex
		lib.Concurrently(urls, lib.DownloadURL)
	})
}
