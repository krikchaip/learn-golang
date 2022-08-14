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
		// for _, url := range urls {
		// 	lib.DownloadURL(url)
		// }

		lib.Concurrently(urls, lib.DownloadURL)
	})
}
