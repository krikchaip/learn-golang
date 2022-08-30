package main

import (
	"21-acceptance-testing/lib"
	"log"
	"net/http"
)

func main() {
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(lib.SlowHandler),
	}

	server := lib.WrapServer(httpServer)

	if err := server.ListenAndServe(); err != nil {
		// this will typically happen if our responses aren't
		// written before the ctx deadline, not much can be done
		log.Fatalf("uh oh, didnt shutdown gracefully, some responses may have been lost %v", err)
	}

	// hopefully, you'll always see this instead
	log.Println("shutdown gracefully! all responses were sent")
}
