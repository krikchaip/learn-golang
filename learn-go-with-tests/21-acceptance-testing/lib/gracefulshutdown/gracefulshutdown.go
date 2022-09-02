package gracefulshutdown

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

type (
	HttpServer interface {
		ListenAndServe() error
		Shutdown(ctx context.Context) error
	}

	HttpServerOption func(*decoratedServer)

	// internal implementation
	decoratedServer struct {
		server HttpServer
		signal <-chan os.Signal
	}
)

// returns a decorated server with graceful shutdown
func NewServer(s HttpServer, options ...HttpServerOption) HttpServer {
	ds := &decoratedServer{
		server: s,
		signal: newSignal(),
	}

	// apply each option to server
	for _, fn := range options {
		fn(ds)
	}

	return ds
}

// allows you to listen to whatever signals you like, rather than the default ones.
func WithSignal(sig <-chan os.Signal) HttpServerOption {
	return func(ds *decoratedServer) {
		ds.signal = sig
	}
}

func (ds *decoratedServer) ListenAndServe() error {
	errChan := make(chan error)

	go func() {
		if err := ds.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
		}
	}()

	go func() {
		// wait for interrupt signals
		fmt.Println("signal received:", <-ds.signal)

		ctx := context.TODO()

		if err := ds.server.Shutdown(ctx); err != nil {
			errChan <- err
		}
	}()

	return <-errChan
}

func (ds *decoratedServer) Shutdown(ctx context.Context) error {
	return ds.server.Shutdown(ctx)
}

func newSignal() <-chan os.Signal {
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, os.Kill, os.Interrupt)
	return sig
}
