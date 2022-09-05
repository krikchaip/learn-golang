package server_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	server "22-building-application/server"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	}}
	sv := server.NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		req := newScoreRequest("Pepper")
		res := httptest.NewRecorder()

		sv.ServeHTTP(res, req)

		assertResponseBody(t, res.Body, "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		req := newScoreRequest("Floyd")
		res := httptest.NewRecorder()

		sv.ServeHTTP(res, req)

		assertResponseBody(t, res.Body, "10")
	})
}

// implements: server.PlayerStore
type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
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
