package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	fp "path/filepath"
	"slices"

	"gopkg.in/yaml.v3"
)

const ADDR = ":8000"

var FILE string

func init() {
	flag.StringVar(&FILE, "file", "", "an optional .yaml file for YAMLHandler")
	flag.Parse()
}

func main() {
	URLShortener := YAMLHandlerFile(FILE).then(YAMLHandlerBlob([]byte(`
    - path: /urlshort
      url: https://github.com/gophercises/urlshort
    - path: /urlshort-final
      url: https://github.com/gophercises/urlshort/tree/solution
  `))).then(MapHandler(map[string]string{
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

func YAMLHandlerBlob(blob []byte) Middleware {
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

func YAMLHandlerFile(file string) Middleware {
	var urlPaths []URLPathYAML

	return func(next http.Handler) http.Handler {
		// skip this middleware if no file is present
		if file == "" {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				next.ServeHTTP(w, r)
			})
		}

		filepath, err := fp.Abs(file)
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Open(filepath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		decoder := yaml.NewDecoder(f)
		if err := decoder.Decode(&urlPaths); err != nil {
			log.Fatal(err)
		}

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
