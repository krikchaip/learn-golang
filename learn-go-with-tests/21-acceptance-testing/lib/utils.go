package lib

import (
	"log"
	"net/http"
	"time"
)

func SlowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Minute)
	w.Write([]byte("Hello World!"))
}

func catch() {
	if err := recover(); err != nil {
		log.Fatal(err)
	}
}
