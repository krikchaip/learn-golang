package gracefulshutdown

import (
	"os"
	"os/signal"
)

type (
	HttpServer interface {
		ListenAndServe() error
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
	// go func() {
	// 	fmt.Printf("%+v\n", <-sig)
	// 	sv.Shutdown(context.TODO())
	// }()

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
	return ds.server.ListenAndServe()
}

func newSignal() <-chan os.Signal {
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, os.Kill, os.Interrupt)
	return sig
}
