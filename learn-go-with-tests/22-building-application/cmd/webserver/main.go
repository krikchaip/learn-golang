package main

import (
	"log"
	"net/http"

	"22-building-application/controller/server"
	"22-building-application/entity"
	"22-building-application/store"
)

const DB_SOURCE = "assets/game.db.json"

// go run 22-building-application/cmd/webserver OR
// cd 22-building-application/cmd/webserver && go run .
func main() {
	st, close := store.SetupFileSystemStore()
	defer close()

	game := entity.NewTexasHoldem(entity.Alerter, st)
	sv := server.NewPlayerServer(st, game)

	// we wrap the call in log.Fatal
	// just in case if there is a problem with ListenAndServe.
	// eg. port already being used, etc.
	log.Fatal(http.ListenAndServe(":3000", sv))
}
