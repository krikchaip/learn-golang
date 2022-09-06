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

		assertStatus(t, res.Code, http.StatusOK)
		assertResponseBody(t, res.Body, "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		req := newScoreRequest("Floyd")
		res := httptest.NewRecorder()

		sv.ServeHTTP(res, req)

		assertStatus(t, res.Code, http.StatusOK)
		assertResponseBody(t, res.Body, "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		req := newScoreRequest("Apollo")
		res := httptest.NewRecorder()

		sv.ServeHTTP(res, req)

		assertStatus(t, res.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{map[string]int{}}
	sv := server.NewPlayerServer(&store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/players/Pepper", nil)
		res := httptest.NewRecorder()

		sv.ServeHTTP(res, req)

		assertStatus(t, res.Code, http.StatusAccepted)
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

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got *bytes.Buffer, want string) {
	t.Helper()
	if got.String() != want {
		t.Errorf("got %q want %q", got, want)
	}
}
