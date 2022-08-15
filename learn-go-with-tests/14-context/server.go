package context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel() // to cancel any operation from Store
}

// an HTTP handler factory which responds to the results of store.Fetch()
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data := make(chan string)

		go func() {
			// fetching data could take some time
			data <- store.Fetch()
		}()

		select {
		// if the user somehow cancels this request, stop Fetching
		case <-ctx.Done():
			store.Cancel()
		case value := <-data:
			fmt.Fprint(w, value)
		}
	}
}
