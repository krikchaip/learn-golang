package main

import (
	"flag"
	"log"
	"net/http"
)

// follows the pattern of "host:port"
// the absence of host means that the handler will listen to every host requested
// on the specific port
// const PORT = ":4000"

func main() {
	// define a command-line flag called "addr"
	addr := flag.String("addr", ":4000", "HTTP network address")

	// NOTE: Must be called after all flags are defined and before flags are accessed
	flag.Parse()

	// flag value is actally a pointer
	PORT := *addr

	router := http.NewServeMux()

	// serve files out of the "./ui/static" directory
	fileServer := http.FileServer(http.Dir("ui/static"))

	// match a single slash, followed by nothing else (exact match)
	router.HandleFunc("GET /{$}", home)

	// this will match the specified pattern exactly
	router.HandleFunc("GET /snippet/view/{id}", snippetView)

	router.HandleFunc("GET /snippet/create", snippetCreate)
	router.HandleFunc("POST /snippet/create", snippetCreatePost)

	// serves static files (subtree path pattern)
	// will match "/static/**", eg. "/static/css/main.css"
	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// a catch-all handler (subtree path pattern)
	// will match "/**", eg. "/foo", "/bar/bax/..."
	// router.HandleFunc("/", defaultHandler)

	log.Printf("starting server on %s", PORT)

	if err := http.ListenAndServe(PORT, router); err != nil {
		log.Fatal(err)
	}
}
