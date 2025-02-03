package main

import (
	"fmt"
	"net/http"
)

const ADDR = ":8000"

func main() {
	URLShortener := YAMLHandler(nil)(MapHandler(nil)(FallbackHandler))

	http.ListenAndServe(ADDR, URLShortener)
}

var FallbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
})

func MapHandler(urlPaths map[string]string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("hello from MapHandler")
			next.ServeHTTP(w, r)
		})
	}
}

func YAMLHandler(yaml []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("hello from YAMLHandler")
			next.ServeHTTP(w, r)
		})
	}
}
