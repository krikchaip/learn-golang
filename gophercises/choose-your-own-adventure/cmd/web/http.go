package main

import (
	"log"
	"net/http"
)

func serveHTTP(addr string) {
	log.Println("Server running at", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
