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
		// v1(store, w, r)
		v2(store, w, r)
	}
}

// handle cancellation logic inside a request handler
func v1(store Store, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := make(chan string)

	go func() {
		// fetching data could take some time
		val, _ := store.Fetch(ctx)
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

// let "store" handle the cancellation by itself (simply pass r.Context)
func v2(store Store, w http.ResponseWriter, r *http.Request) {
	data, err := store.Fetch(r.Context())

	// TODO: log error however you like
	if err != nil {
		return
	}

	fmt.Fprint(w, data)
}
