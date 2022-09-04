package server

import "net/http"

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("20"))
}
