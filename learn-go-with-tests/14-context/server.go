package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel() // to cancel any operation from Store
}

// an HTTP handler factory which responds to the results of store.Fetch()
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v1(store, w, r)
	}
}

func v1(store Store, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := make(chan string)

	go func() {
		// fetching data could take some time
		val, _ := store.Fetch(nil)
		data <- val
	}()

	select {
	// if the user somehow cancels this request, stop Fetching
	case <-ctx.Done():
		store.Cancel()
	case value := <-data:
		fmt.Fprint(w, value)
	}
}
