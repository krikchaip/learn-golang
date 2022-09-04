package main

import (
	"22-building-application/server"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)
	log.Fatal(http.ListenAndServe(":3000", handler))
}
