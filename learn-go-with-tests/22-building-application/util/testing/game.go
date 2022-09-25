package testing

// implements: entity.Game
type GameSpy struct {
	StartedWith  int
	FinishedWith string
}

func (g *GameSpy) Start(nPlayers int) {
	g.StartedWith = nPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}
