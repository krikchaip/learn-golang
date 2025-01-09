package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	// compose a middleware chain using 'alice' package
	// which will be used for every request our application receives.
	standard := alice.New(app.recoverPanic, app.logRequest, securityHeaders)

	// unprotected application routes using the "dynamic" middleware chain.
	dynamic := alice.New(app.sessionManager.LoadAndSave)

	// a middleware that allows only for authenticated requests
	protected := dynamic.Append(app.requireAuthentication)

	// serve files out of the "./ui/static" directory
	fileServer := http.FileServer(http.Dir("ui/static"))

	// serves static files (subtree path pattern)
	// will match "/static/**", eg. "/static/css/main.css"
	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// a catch-all handler (subtree path pattern)
	// will match "/**", eg. "/foo", "/bar/bax/..."
	// router.HandleFunc("/", app.defaultHandler)

	// match a single slash, followed by nothing else (exact match)
	router.Handle("GET /{$}", dynamic.ThenFunc(app.home))

	// this will match the specified pattern exactly
	router.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))

	router.Handle("GET /snippet/create", protected.ThenFunc(app.snippetCreate))
	router.Handle("POST /snippet/create", protected.ThenFunc(app.snippetCreatePost))

	router.Handle("GET /user/signup", dynamic.ThenFunc(app.userSignup))
	router.Handle("POST /user/signup", dynamic.ThenFunc(app.userSignupPost))

	router.Handle("GET /user/login", dynamic.ThenFunc(app.userLogin))
	router.Handle("POST /user/login", dynamic.ThenFunc(app.userLoginPost))

	router.Handle("POST /user/logout", protected.ThenFunc(app.userLogoutPost))

	// wrap ServeMux router with middlewares.
	// do note that ServeMux also implements the 'http.Handler' interface
	return standard.Then(router)
}
