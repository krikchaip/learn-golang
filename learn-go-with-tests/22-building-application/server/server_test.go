package server_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	server "22-building-application/server"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		req := newScoreRequest("Pepper")
		res := httptest.NewRecorder()

		server.PlayerServer(res, req)

		assertResponseBody(t, res.Body, "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		req := newScoreRequest("Floyd")
		res := httptest.NewRecorder()

		server.PlayerServer(res, req)

		assertResponseBody(t, res.Body, "10")
	})
}

func newScoreRequest(name string) *http.Request {
	return httptest.NewRequest(http.MethodGet, "/players/"+name, nil)
}

func assertResponseBody(t testing.TB, got *bytes.Buffer, want string) {
	t.Helper()
	if got.String() != want {
		t.Errorf("got %q want %q", got, want)
	}
}
