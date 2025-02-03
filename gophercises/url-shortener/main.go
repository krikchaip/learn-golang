package main

import (
	"fmt"
	"log"
	"net/http"
)

const ADDR = ":8000"

func main() {
	URLShortener := MapHandler(map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		"/github":         "https://github.com/krikchaip",
	})

	http.ListenAndServe(ADDR, URLShortener(FallbackHandler))
}

var FallbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
})

func MapHandler(urlPaths map[string]string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if url, ok := urlPaths[r.URL.Path]; ok {
				log.Println(r.URL.Path, "->", url)
				http.Redirect(w, r, url, http.StatusMovedPermanently)

				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func YAMLHandler(yaml []byte) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("hello from YAMLHandler")
			next.ServeHTTP(w, r)
		})
	}
}
