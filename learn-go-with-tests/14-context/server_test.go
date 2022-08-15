package context_test

import (
	lib "14-context"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	t.Run("handler responds with a mocked response", func(t *testing.T) {
		mockedResponse := "Hello, World!"
		store := &StoreSpy{
			response: mockedResponse,
			t:        t,
		}
		handler := lib.Server(store)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		handler.ServeHTTP(w, r)

		store.assertNotCancelled()

		if w.Body.String() != mockedResponse {
			t.Errorf("got %q, want %q", w.Body.String(), mockedResponse)
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		mockedResponse := "Hello, World!"
		store := &StoreSpy{
			response: mockedResponse,
			delay:    100 * time.Millisecond,
			t:        t,
		}
		handler := lib.Server(store)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		r = cancelWithin(r, 20*time.Millisecond)
		handler.ServeHTTP(w, r)

		store.assertCancelled()
	})
}

type StoreSpy struct {
	response string // mocked response
	delay    time.Duration
	canceled bool
	t        testing.TB
}

func (s *StoreSpy) Fetch() string {
	time.Sleep(s.delay) // give some time for the user to cancel
	return s.response
}

func (s *StoreSpy) Cancel() {
	s.canceled = true
}

func (s *StoreSpy) assertCancelled() {
	s.t.Helper()
	if !s.canceled {
		s.t.Error("store was not told to cancel")
	}
}

func (s *StoreSpy) assertNotCancelled() {
	s.t.Helper()
	if s.canceled {
		s.t.Error("it should not have cancelled the store")
	}
}

func cancelWithin(r *http.Request, t time.Duration) *http.Request {
	newContext, cancel := context.WithCancel(r.Context())
	time.AfterFunc(t, cancel)
	return r.Clone(newContext)
}
