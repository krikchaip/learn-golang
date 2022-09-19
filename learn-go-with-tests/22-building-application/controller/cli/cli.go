package cli

import (
	"22-building-application/entity"
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

const PlayerPrompt = "Please enter the number of players: "

type CLI struct {
	store   entity.PlayerStore
	scanner *bufio.Scanner
	printer io.Writer
	alerter entity.BlindAlerter
}

func NewPlayerCLI(
	store entity.PlayerStore,
	in io.Reader,
	out io.Writer,
	alerter entity.BlindAlerter,
) *CLI {
	return &CLI{
		store,
		bufio.NewScanner(in),
		out,
		alerter,
	}
}

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.printer, PlayerPrompt)

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

	if len(xs) != 2 {
		ok = false
		return
	}

	name, cmd := xs[0], xs[1]

	if cmd != "wins" {
		ok = false
		return
	}

	winner = name
	ok = true

	return
}
