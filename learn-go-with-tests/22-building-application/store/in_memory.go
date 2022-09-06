package store

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		store: make(map[string]int),
	}
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return s.store[name]
}

func (s *InMemoryPlayerStore) RecordWin(name string) {
	s.store[name]++
}
