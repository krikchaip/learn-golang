package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"22-building-application/server"
	"22-building-application/store"
	util "22-building-application/util/testing"
)

// integration testing of server.PlayerServer & store.InMemoryPlayerStore
func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	st := store.NewInMemoryPlayerStore()
	sv := server.NewPlayerServer(st)

	player := "Pepper"

	sv.ServeHTTP(httptest.NewRecorder(), util.NewPostWinRequest(player))
	sv.ServeHTTP(httptest.NewRecorder(), util.NewPostWinRequest(player))
	sv.ServeHTTP(httptest.NewRecorder(), util.NewPostWinRequest(player))

	res := httptest.NewRecorder()
	sv.ServeHTTP(res, util.NewScoreRequest(player))

	util.AssertStatus(t, res.Code, http.StatusOK)
	util.AssertResponseBody(t, res.Body, "3")
}
