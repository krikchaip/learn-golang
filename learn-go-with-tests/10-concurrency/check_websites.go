package concurrency

// returns whether the website returned a good or a bad response
type WebsiteChecker func(url string) bool

type result struct {
	string // url
	bool   // response status
}

// ?? run this program with `go test -race` to debug race condition
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		// ?? run this anonymous function in a separate process
		go func(u string) {
			// ?? send statement (await receiver, "blocking call")
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for range urls {
		// ?? receive expression (await sender, "blocking call")
		r := <-resultChannel

		results[r.string] = r.bool
	}

	return results
}
