package main

import (
	"fmt"
	"net/http"
)

const ADDR = ":8000"

func main() {
	URLShortener := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	http.ListenAndServe(ADDR, URLShortener)
}
