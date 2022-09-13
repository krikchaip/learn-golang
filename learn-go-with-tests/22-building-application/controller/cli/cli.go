package cli

import (
	"22-building-application/entity"
	"bufio"
	"io"
	"strings"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

type CLI struct {
	store   entity.PlayerStore
	scanner *bufio.Scanner
	alerter BlindAlerter
}

func NewPlayerCLI(
	store entity.PlayerStore,
	in io.Reader,
	alerter BlindAlerter,
) *CLI {
	return &CLI{
		store,
		bufio.NewScanner(in),
		alerter,
	}
}

func (c *CLI) PlayPoker() {
	c.scheduleBlindAlert()

	input := c.readLine()
	name, ok := extractWinner(input)

	if ok {
		c.store.RecordWin(name)
	}
}

func (c *CLI) scheduleBlindAlert() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, b := range blinds {
		c.alerter.ScheduleAlertAt(blindTime, b)
		blindTime = blindTime + 10*time.Minute
	}
}

func (c *CLI) readLine() string {
	c.scanner.Scan()
	return c.scanner.Text()
}

func extractWinner(input string) (winner string, ok bool) {
	xs := strings.SplitN(input, " ", 2)
	name, cmd := xs[0], xs[1]

	if cmd != "wins" {
		ok = false
		return
	}

	winner = name
	ok = true

	return
}
