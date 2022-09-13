package cli_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"22-building-application/controller/cli"
	testutil "22-building-application/util/testing"
)

var (
	dummySpyAlerter = &testutil.SpyBlindAlerter{}
)

func TestCLI(t *testing.T) {
	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		store := testutil.NewStubPlayerStore()

		program := cli.NewPlayerCLI(store, in, dummySpyAlerter)
		program.PlayPoker()

		testutil.AssertPlayerWin(t, store.GetWinCalls(), []string{"Chris"})
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		store := testutil.NewStubPlayerStore()

		program := cli.NewPlayerCLI(store, in, dummySpyAlerter)
		program.PlayPoker()

		testutil.AssertPlayerWin(t, store.GetWinCalls(), []string{"Cleo"})
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		store := testutil.NewStubPlayerStore()
		blindAlerter := &testutil.SpyBlindAlerter{}

		program := cli.NewPlayerCLI(store, in, blindAlerter)
		program.PlayPoker()

		cases := []testutil.ScheduleAlert{
			{Duration: 0 * time.Second, Amount: 100},
			{Duration: 10 * time.Minute, Amount: 200},
			{Duration: 20 * time.Minute, Amount: 300},
			{Duration: 30 * time.Minute, Amount: 400},
			{Duration: 40 * time.Minute, Amount: 500},
			{Duration: 50 * time.Minute, Amount: 600},
			{Duration: 60 * time.Minute, Amount: 800},
			{Duration: 70 * time.Minute, Amount: 1000},
			{Duration: 80 * time.Minute, Amount: 2000},
			{Duration: 90 * time.Minute, Amount: 4000},
			{Duration: 100 * time.Minute, Amount: 8000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				// this should always greater than the index
				if len(blindAlerter.Alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
				}

				got := blindAlerter.Alerts[i]
				testutil.AssertScheduledAlert(t, got, want)
			})
		}
	})
}
