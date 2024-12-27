package main

import (
	"log"
	"net/http"
)

// follows the pattern of "host:port"
// the absence of host means that the handler will listen to every host requested
// on the specific port
const PORT = ":4000"

func main() {
	router := http.NewServeMux()

	// match a single slash, followed by nothing else (exact match)
	router.HandleFunc("GET /{$}", home)

	// this will match the specified pattern exactly
	router.HandleFunc("GET /snippet/view/{id}", snippetView)

	router.HandleFunc("GET /snippet/create", snippetCreate)
	router.HandleFunc("POST /snippet/create", snippetCreatePost)

	// a catch-all handler (subtree path pattern)
	// will match "/**", eg. "/foo", "/bar/bax/..."
	// router.HandleFunc("/", defaultHandler)

	log.Printf("starting server on %s", PORT)

	if err := http.ListenAndServe(PORT, router); err != nil {
		log.Fatal(err)
	}
}
