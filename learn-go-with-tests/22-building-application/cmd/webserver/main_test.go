package main_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"

	"22-building-application/controller/server"
	"22-building-application/entity"
	"22-building-application/store"
	util "22-building-application/util/testing"
)

// integration testing of server.PlayerServer & store.InMemoryPlayerStore
func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	st := setupStore(t)
	game := entity.NewTexasHoldem(entity.Alerter, st)
	sv := server.NewPlayerServer(st, game)

	player := "Pepper"

	sv.ServeHTTP(httptest.NewRecorder(), util.NewPostWinRequest(player))
	sv.ServeHTTP(httptest.NewRecorder(), util.NewPostWinRequest(player))
	sv.ServeHTTP(httptest.NewRecorder(), util.NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		res := httptest.NewRecorder()
		sv.ServeHTTP(res, util.NewScoreRequest(player))

		util.AssertStatus(t, res.Code, http.StatusOK)
		util.AssertResponseBody(t, res.Body, "3")
	})

	t.Run("get league", func(t *testing.T) {
		res := httptest.NewRecorder()
		sv.ServeHTTP(res, util.NewLeagueRequest())

		got := util.ParseLeagueFromResponse(t, res.Body)
		want := []entity.Player{
			{Name: player, Wins: 3},
		}

		util.AssertStatus(t, res.Code, http.StatusOK)
		util.AssertLeagueTable(t, got, want)
	})
}

func TestConcurrentRecordingWins(t *testing.T) {
	st := setupStore(t)
	game := entity.NewTexasHoldem(entity.Alerter, st)
	sv := server.NewPlayerServer(st, game)

	player := "Pepper"
	nConcurrent := 1000

	var wg sync.WaitGroup
	wg.Add(nConcurrent)

	for i := 0; i < nConcurrent; i++ {
		go func() {
			sv.ServeHTTP(httptest.NewRecorder(), util.NewPostWinRequest(player))
			wg.Done()
		}()
	}

	wg.Wait()

	res := httptest.NewRecorder()
	sv.ServeHTTP(res, util.NewScoreRequest(player))

	util.AssertStatus(t, res.Code, http.StatusOK)
	util.AssertResponseBody(t, res.Body, strconv.Itoa(nConcurrent))
}

func TestEmptyFileSource(t *testing.T) {
	src, cleanup := util.CreateTempFile(t, "")
	t.Cleanup(cleanup)

	util.AssertNoPanic(t, func() {
		store.NewFileSystemPlayerStore(src)
	})
}

func setupStore(t testing.TB) entity.PlayerStore {
	t.Helper()

	// st := store.NewInMemoryPlayerStore()

	src, cleanup := util.CreateTempFile(t, "[]")
	st := store.NewFileSystemPlayerStore(src)
	t.Cleanup(cleanup)

	return st
}
