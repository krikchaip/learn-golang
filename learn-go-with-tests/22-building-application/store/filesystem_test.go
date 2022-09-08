package store_test

import (
	"22-building-application/server"
	"22-building-application/store"
	"22-building-application/util/testing"
	"strings"
	tt "testing"
)

func TestFileSystemStore(t *tt.T) {
	t.Run("league from a reader", func(t *tt.T) {
		src := strings.NewReader(`[
			{ "Name": "Cleo",  "Wins": 10 },
			{ "Name": "Chris", "Wins": 33 }
		]`)
		store := store.NewFileSystemPlayerStore(src)

		got := store.GetLeagueTable()
		want := []server.Player{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 33},
		}

		testing.AssertLeagueTable(t, got, want)
	})
}
