package store_test

import (
	"os"
	tt "testing"

	"22-building-application/entity"
	"22-building-application/store"
	"22-building-application/util/testing"
)

func TestFileSystemStore(t *tt.T) {
	t.Run("league from a reader", func(t *tt.T) {
		src := setupSource(t)
		store := store.NewFileSystemPlayerStore(src)

		got := store.GetLeagueTable()
		want := []entity.Player{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 33},
		}

		testing.AssertLeagueTable(t, got, want)
	})

	t.Run("return the same result on second call", func(t *tt.T) {
		src := setupSource(t)
		store := store.NewFileSystemPlayerStore(src)

		// first time calling
		store.GetLeagueTable()

		got := store.GetLeagueTable()
		want := []entity.Player{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 33},
		}

		testing.AssertLeagueTable(t, got, want)
	})

	t.Run("get player score", func(t *tt.T) {
		src := setupSource(t)
		store := store.NewFileSystemPlayerStore(src)

		got := store.GetPlayerScore("Chris")
		want := 33

		testing.AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *tt.T) {
		src := setupSource(t)
		store := store.NewFileSystemPlayerStore(src)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34

		testing.AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *tt.T) {
		src := setupSource(t)
		store := store.NewFileSystemPlayerStore(src)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1

		testing.AssertScoreEquals(t, got, want)
	})
}

func setupSource(t tt.TB) (src *os.File) {
	t.Helper()

	// // ?? does not implement io.Writer() (only io.Reader, io.Seeker)
	// src := strings.NewReader(`[
	// 	{ "Name": "Cleo",  "Wins": 10 },
	// 	{ "Name": "Chris", "Wins": 33 }
	// ]`)

	// ?? os.File implements io.ReadWriteSeeker
	src, cleanup := testing.CreateTempFile(t, `[
		{ "Name": "Cleo",  "Wins": 10 },
		{ "Name": "Chris", "Wins": 33 }
	]`)

	// ** don't forget to cleanup after tests
	t.Cleanup(cleanup)

	return
}
