package store_test

import (
	tt "testing"

	"22-building-application/server"
	"22-building-application/store"
	"22-building-application/util/testing"
)

func TestFileSystemStore(t *tt.T) {
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

	t.Run("league from a reader", func(t *tt.T) {
		store := store.NewFileSystemPlayerStore(src)

		got := store.GetLeagueTable()
		want := []server.Player{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 33},
		}

		testing.AssertLeagueTable(t, got, want)
	})

	t.Run("return the same result on second call", func(t *tt.T) {
		store := store.NewFileSystemPlayerStore(src)

		// first time calling
		store.GetLeagueTable()

		got := store.GetLeagueTable()
		want := []server.Player{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 33},
		}

		testing.AssertLeagueTable(t, got, want)
	})

	t.Run("get player score", func(t *tt.T) {
		store := store.NewFileSystemPlayerStore(src)

		got := store.GetPlayerScore("Chris")
		want := 33

		testing.AssertScoreEquals(t, got, want)
	})
}