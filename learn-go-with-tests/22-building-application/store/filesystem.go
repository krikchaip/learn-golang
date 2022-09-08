package store

import (
	"22-building-application/server"
	"io"
)

// implements: server.PlayerStore
type FileSystemPlayerStore struct {
	source io.Reader
}

func NewFileSystemPlayerStore(source io.Reader) *FileSystemPlayerStore {
	return &FileSystemPlayerStore{source}
}

func (s *FileSystemPlayerStore) GetLeagueTable() (league []server.Player) {
	league, _ = server.NewLeague(s.source)
	return
}
