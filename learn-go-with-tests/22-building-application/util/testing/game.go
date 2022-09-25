package testing

// implements: entity.Game
type GameSpy struct {
	StartCalled bool
	StartedWith int

	FinishCalled bool
	FinishedWith string
}

func (g *GameSpy) Start(nPlayers int) {
	g.StartCalled = true
	g.StartedWith = nPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishCalled = true
	g.FinishedWith = winner
}