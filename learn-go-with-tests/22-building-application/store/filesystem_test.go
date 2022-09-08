package store_test

import (
	"strings"
	tt "testing"

	"22-building-application/server"
	"22-building-application/store"
	"22-building-application/util/testing"
)

func TestFileSystemStore(t *tt.T) {
	src := strings.NewReader(`[
		{ "Name": "Cleo",  "Wins": 10 },
		{ "Name": "Chris", "Wins": 33 }
	]`)

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
