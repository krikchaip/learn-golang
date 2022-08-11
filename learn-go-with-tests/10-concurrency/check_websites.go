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
			resultChannel <- result{u, wc(u)} // ?? send statement
		}(url)
	}

	for range urls {
		r := <-resultChannel // ?? receive expression
		results[r.string] = r.bool
	}

	return results
}
