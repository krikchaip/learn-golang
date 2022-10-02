package cli

import (
	"22-building-application/entity"
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	PlayerPrompt         = "Please enter the number of players: "
	BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"
)

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game entity.Game
}

func NewPlayerCLI(
	in io.Reader,
	out io.Writer,
	game entity.Game,
) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)

	nPlayers, err := strconv.Atoi(c.readLine())
	if err != nil {
		fmt.Fprint(c.out, BadPlayerInputErrMsg)
		return
	}

	c.game.Start(c.out, nPlayers)

	input := c.readLine()
	name, ok := extractWinner(input)

	if ok {
		c.game.Finish(name)
	}
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
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
