package main

import "net/http"

type Middleware func(http.Handler) http.Handler

func (m Middleware) then(next Middleware) Middleware {
	return func(h http.Handler) http.Handler {
		return m(next(h))
	}
}
