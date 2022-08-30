package lib

import (
	"context"
	"net/http"
	"os"
	"os/signal"
)

// returns a decorated server with graceful shutdown
func WrapServer(sv *http.Server) *http.Server {
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, os.Kill, os.Interrupt)

	go func() {
		<-sig
		sv.Shutdown(context.TODO())
	}()

	return sv
}
