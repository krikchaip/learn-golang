package store

import (
	"22-building-application/server"
	"io"
)

// implements: server.PlayerStore
type FileSystemPlayerStore struct {
	source io.ReadSeeker
}

func NewFileSystemPlayerStore(source io.ReadSeeker) *FileSystemPlayerStore {
	return &FileSystemPlayerStore{source}
}

func (s *FileSystemPlayerStore) GetLeagueTable() (league []server.Player) {
	s.source.Seek(0, io.SeekStart)
	league, _ = server.NewLeague(s.source)
	return
}
