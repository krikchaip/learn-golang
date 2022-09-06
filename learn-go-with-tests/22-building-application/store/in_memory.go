package store

import "sync"

type InMemoryPlayerStore struct {
	store map[string]int
	mut   sync.Mutex
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
	s.mut.Lock()
	defer s.mut.Unlock()

	s.store[name]++
}
