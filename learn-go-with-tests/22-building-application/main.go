package main

import (
	"log"
	"net/http"

	"22-building-application/server"
	"22-building-application/store"
)

func main() {
	sv := server.NewPlayerServer(&store.InMemoryPlayerStore{})

	// we wrap the call in log.Fatal
	// just in case if there is a problem with ListenAndServe.
	// eg. port already being used, etc.
	log.Fatal(http.ListenAndServe(":3000", sv))
}
