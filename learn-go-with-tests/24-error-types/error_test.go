package error_types_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	error_types "24-error-types"
)

func TestDumbGetter(t *testing.T) {
	t.Run("when you don't get a 200 you get a status error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusForbidden)
		}))
		defer server.Close()

		_, err := error_types.DumbGetter(server.URL)
		if err == nil {
			t.Fatal("expected an error")
		}

		// ?? assert for the error type instead of the error messages
		// got, ok := err.(error_types.BadStatusError)

		// ** a modern alternative
		var got error_types.BadStatusError
		ok := errors.As(err, &got)

		if !ok {
			t.Fatalf("was not a BadStatusError, got %T", err)
		}

		want := error_types.BadStatusError{
			URL:    server.URL,
			Status: http.StatusForbidden,
		}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
