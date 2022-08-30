package lib

import (
	"net/http"
	"time"
)

func SlowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Minute)
	w.Write([]byte("Hello World!"))
}
