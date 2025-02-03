package main

import (
	"fmt"
	"log"
	"net/http"
	"slices"

	"gopkg.in/yaml.v3"
)

const ADDR = ":8000"

func main() {
	URLShortener := YAMLHandler([]byte(`
    - path: /urlshort
      url: https://github.com/gophercises/urlshort
    - path: /urlshort-final
      url: https://github.com/gophercises/urlshort/tree/solution
  `)).then(MapHandler(map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		"/github":         "https://github.com/krikchaip",
	}))

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

type URLPathYAML struct {
	Path string
	Url  string
}

func YAMLHandler(blob []byte) Middleware {
	var urlPaths []URLPathYAML

	if err := yaml.Unmarshal(blob, &urlPaths); err != nil {
		log.Fatal(err)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			i := slices.IndexFunc(urlPaths, func(rec URLPathYAML) bool {
				return rec.Path == r.URL.Path
			})

			if i != -1 {
				url := urlPaths[i].Url

				log.Println(r.URL.Path, "->", url)
				http.Redirect(w, r, url, http.StatusMovedPermanently)

				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
