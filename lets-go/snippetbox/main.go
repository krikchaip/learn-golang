package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// follows the pattern of "host:port"
// the absence of host means that the handler will listen to every host requested
// on the specific port
const PORT = ":4000"

func home(w http.ResponseWriter, r *http.Request) {
	// add a new header key
	// NOTE: MUST BE executed before any call to WriteHeader() or Write()
	w.Header().Add("X-App-Env", "development")

	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	// retrieving path param
	id, err := strconv.Atoi(r.PathValue("id"))

	// field validation
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
}

func snippetCreate(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// send a 201 status code.
	// NOTE: MUST BE executed before any call to Write()
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Save a new snippet..."))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %q!", r.URL.Path)
}

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
