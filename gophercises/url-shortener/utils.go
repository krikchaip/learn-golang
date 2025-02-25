package main

import "net/http"

type Middleware func(http.Handler) http.Handler

func (m Middleware) then(next Middleware) Middleware {
	return func(h http.Handler) http.Handler {
		return m(next(h))
	}
}

var NoopMiddleware Middleware = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
