package error_types_test

import (
	"fmt"
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

		got := err.Error()
		want := fmt.Sprintf(
			"did not get 200 from %s, got %d",
			server.URL,
			http.StatusForbidden,
		)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
