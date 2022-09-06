package main

import (
	"22-building-application/server"
	"log"
	"net/http"
)

func main() {
	sv := server.NewPlayerServer(&InMemoryPlayerStore{})

	// we wrap the call in log.Fatal
	// just in case if there is a problem with ListenAndServe.
	// eg. port already being used, etc.
	log.Fatal(http.ListenAndServe(":3000", sv))
}

// TODO: will implement later
type InMemoryPlayerStore struct{}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 0
}

func (s *InMemoryPlayerStore) RecordWin(name string) {}
