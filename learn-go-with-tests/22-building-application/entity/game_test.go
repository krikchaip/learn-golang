package entity_test

import (
	"fmt"
	"testing"
	"time"

	"22-building-application/entity"
	testutil "22-building-application/util/testing"
)

var (
	dummyPlayerStore  = testutil.NewStubPlayerStore()
	dummyBlindAlerter = &testutil.SpyBlindAlerter{}
)

func TestGame_Start(t *testing.T) {
	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		alerter := &testutil.SpyBlindAlerter{}
		game := entity.NewGame(alerter, dummyPlayerStore)

		game.Start(5)

		cases := []testutil.ScheduleAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 10 * time.Minute, Amount: 200},
			{At: 20 * time.Minute, Amount: 300},
			{At: 30 * time.Minute, Amount: 400},
			{At: 40 * time.Minute, Amount: 500},
			{At: 50 * time.Minute, Amount: 600},
			{At: 60 * time.Minute, Amount: 800},
			{At: 70 * time.Minute, Amount: 1000},
			{At: 80 * time.Minute, Amount: 2000},
			{At: 90 * time.Minute, Amount: 4000},
			{At: 100 * time.Minute, Amount: 8000},
		}

		checkSchedulingCases(t, cases, alerter)
	})

	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		alerter := &testutil.SpyBlindAlerter{}
		game := entity.NewGame(alerter, dummyPlayerStore)

		game.Start(7)

		cases := []testutil.ScheduleAlert{
			{At: 0 * time.Minute, Amount: 100},
			{At: 12 * time.Minute, Amount: 200},
			{At: 24 * time.Minute, Amount: 300},
			{At: 36 * time.Minute, Amount: 400},
		}

		checkSchedulingCases(t, cases, alerter)
	})
}

func TestGame_Finish(t *testing.T) {
	store := testutil.NewStubPlayerStore()
	game := entity.NewGame(dummyBlindAlerter, store)

	winner := "Ruth"
	game.Finish(winner)

	testutil.AssertPlayerWin(t, store.GetWinCalls(), []string{"Ruth"})
}

func checkSchedulingCases(
	t *testing.T,
	cases []testutil.ScheduleAlert,
	alerter *testutil.SpyBlindAlerter,
) {
	t.Helper()
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {
			// this should always greater than the index
			if len(alerter.Alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, alerter.Alerts)
			}

			got := alerter.Alerts[i]
			testutil.AssertScheduledAlert(t, got, want)
		})
	}
}
