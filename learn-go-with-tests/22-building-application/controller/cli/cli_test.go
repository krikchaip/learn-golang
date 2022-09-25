package cli_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"22-building-application/controller/cli"
	testutil "22-building-application/util/testing"
)

func TestCLI(t *testing.T) {
	t.Run("record Chris win from user input", func(t *testing.T) {
		store := testutil.NewStubPlayerStore()
		in := strings.NewReader("5\nChris wins\n")
		out := &bytes.Buffer{}
		blindAlerter := &testutil.SpyBlindAlerter{}

		program := cli.NewPlayerCLI(store, in, out, blindAlerter)
		program.PlayPoker()

		testutil.AssertPlayerWin(t, store.GetWinCalls(), []string{"Chris"})
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		store := testutil.NewStubPlayerStore()
		in := strings.NewReader("5\nCleo wins\n")
		out := &bytes.Buffer{}
		blindAlerter := &testutil.SpyBlindAlerter{}

		program := cli.NewPlayerCLI(store, in, out, blindAlerter)
		program.PlayPoker()

		testutil.AssertPlayerWin(t, store.GetWinCalls(), []string{"Cleo"})
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		store := testutil.NewStubPlayerStore()
		in := strings.NewReader("5\nChris wins\n")
		out := &bytes.Buffer{}
		blindAlerter := &testutil.SpyBlindAlerter{}

		program := cli.NewPlayerCLI(store, in, out, blindAlerter)
		program.PlayPoker()

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

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		store := testutil.NewStubPlayerStore()
		in := strings.NewReader("7\n")
		out := &bytes.Buffer{}
		blindAlerter := &testutil.SpyBlindAlerter{}

		program := cli.NewPlayerCLI(store, in, out, blindAlerter)
		program.PlayPoker()

		got := out.String()
		want := cli.PlayerPrompt

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		cases := []testutil.ScheduleAlert{
			{At: 0 * time.Minute, Amount: 100},
			{At: 12 * time.Minute, Amount: 200},
			{At: 24 * time.Minute, Amount: 300},
			{At: 36 * time.Minute, Amount: 400},
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
