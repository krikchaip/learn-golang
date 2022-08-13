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

	for _, url := range urls {
		lib.DownloadURL(url)
	}
}
