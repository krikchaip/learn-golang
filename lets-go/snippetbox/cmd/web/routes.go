package main

import "net/http"

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	// serve files out of the "./ui/static" directory
	fileServer := http.FileServer(http.Dir("ui/static"))

	// serves static files (subtree path pattern)
	// will match "/static/**", eg. "/static/css/main.css"
	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// a catch-all handler (subtree path pattern)
	// will match "/**", eg. "/foo", "/bar/bax/..."
	// router.HandleFunc("/", app.defaultHandler)

	// match a single slash, followed by nothing else (exact match)
	router.HandleFunc("GET /{$}", app.home)

	// this will match the specified pattern exactly
	router.HandleFunc("GET /snippet/view/{id}", app.snippetView)

	router.HandleFunc("GET /snippet/create", app.snippetCreate)
	router.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	// wrap ServeMux router with the 'securityHeaders' middleware.
	// do note that ServeMux also implements the 'http.Handler' interface
	return securityHeaders(router)
}
