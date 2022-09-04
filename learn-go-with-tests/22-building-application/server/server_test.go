package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	server "22-building-application/server"
)

func TestGETPlayers(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/players/Pepper", nil)
	res := httptest.NewRecorder()

	server.PlayerServer(res, req)

	got := res.Body.String()
	want := "20"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
