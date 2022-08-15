package context_test

import (
	ctx "14-context"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StoreStub struct {
	response string
}

func (s *StoreStub) Fetch() string {
	return s.response
}

func TestServer(t *testing.T) {
	mockedResponse := "Hello, World!"
	handler := ctx.Server(
		&StoreStub{response: mockedResponse},
	)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	handler.ServeHTTP(w, r)

	if w.Body.String() != mockedResponse {
		t.Errorf("got %q, want %q", w.Body.String(), mockedResponse)
	}
}
