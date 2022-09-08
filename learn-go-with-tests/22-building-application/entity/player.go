package entity

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	GetLeagueTable() League
	RecordWin(name string)
}
