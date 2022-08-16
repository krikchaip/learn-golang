package context_test

import (
	lib "14-context"
	"context"
	"fmt"
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

		// httptest.NewRecorder doesn't support the case when
		// you want to know whether a response has been written.
		w := &ResponseWriterSpy{t: t}
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		r = cancelWithin(r, 20*time.Millisecond)
		handler.ServeHTTP(w, r)

		store.assertCancelled()
		w.assertNotWritten()
	})
}

type StoreSpy struct {
	response string // mocked response
	delay    time.Duration
	canceled bool
	t        testing.TB
}

func (s *StoreSpy) Fetch(ctx context.Context) (string, error) {
	data := make(chan string)

	go func() {
		var result string

		// appeding each character one by one with a delay
		// eg. "H...(10ms)...E...(10ms)...L..."
		for _, char := range s.response {
			select {
			case <-ctx.Done(): // stop fetching in the middle
				fmt.Println("spy store got cancelled")
				return
			case <-time.After(s.delay): // give some time for the user to cancel
				result += string(char)
			}
		}

		// puts to "data" after completed
		data <- result
	}()

	select {
	case <-ctx.Done():
		s.Cancel()
		return "", ctx.Err()
	case val := <-data: // await data
		return val, nil
	}
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

// implements: http.ResponseWriter
type ResponseWriterSpy struct {
	written bool // indicates that whether a response has been written
	t       testing.TB
}

func (w *ResponseWriterSpy) Header() http.Header {
	w.written = true
	return nil
}

func (w *ResponseWriterSpy) Write([]byte) (int, error) {
	w.written = true
	return 0, fmt.Errorf("not implemented")
}

func (w *ResponseWriterSpy) WriteHeader(int) {
	w.written = true
}

func (w *ResponseWriterSpy) assertNotWritten() {
	w.t.Helper()
	if w.written {
		w.t.Error("a response should not have been written")
	}
}

func cancelWithin(r *http.Request, t time.Duration) *http.Request {
	newContext, cancel := context.WithCancel(r.Context())
	time.AfterFunc(t, cancel)
	return r.Clone(newContext)
}
