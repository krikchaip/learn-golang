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
		fmt.Fprint(w, store.Fetch())
	}
}
