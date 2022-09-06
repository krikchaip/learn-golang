package store

// TODO: will implement later
type InMemoryPlayerStore struct{}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 0
}

func (s *InMemoryPlayerStore) RecordWin(name string) {}
