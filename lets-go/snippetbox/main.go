package main

import (
	"fmt"
	"log"
	"net/http"
)

// follows the pattern of "host:port"
// the absence of host means that the handler will listen to every host requested
// on the specific port
const PORT = ":4000"

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %q!", r.URL.Path)
}

func main() {
	router := http.NewServeMux()

	// match a single slash, followed by nothing else (exact match)
	router.HandleFunc("/{$}", home)

	// will match the specified pattern exactly
	router.HandleFunc("/snippet/view", snippetView)
	router.HandleFunc("/snippet/create", snippetCreate)

	// a catch-all handler (subtree path pattern)
	// will match "/**", eg. "/foo", "/bar/bax/..."
	// router.HandleFunc("/", defaultHandler)

	log.Printf("starting server on %s", PORT)

	if err := http.ListenAndServe(PORT, router); err != nil {
		log.Fatal(err)
	}
}
