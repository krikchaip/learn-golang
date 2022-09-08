package store

import (
	"22-building-application/server"
	"io"
)

// implements: server.PlayerStore
type FileSystemPlayerStore struct {
	source io.ReadSeeker
}

func NewFileSystemPlayerStore(source io.ReadSeeker) server.PlayerStore {
	return &FileSystemPlayerStore{source}
}

func (s *FileSystemPlayerStore) GetPlayerScore(name string) int {
	for _, p := range s.GetLeagueTable() {
		if name == p.Name {
			return p.Wins
		}
	}

	return 0
}

func (s *FileSystemPlayerStore) GetLeagueTable() (league []server.Player) {
	s.source.Seek(0, io.SeekStart)
	league, _ = server.NewLeague(s.source)
	return
}

func (s *FileSystemPlayerStore) RecordWin(name string) {

}
