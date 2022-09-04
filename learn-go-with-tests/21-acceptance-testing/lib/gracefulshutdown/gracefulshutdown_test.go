package gracefulshutdown_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"21-acceptance-testing/lib/gracefulshutdown"
	"21-acceptance-testing/lib/testutil/assert"
)

func TestGracefulShutdownServer_ListenAndServe(t *testing.T) {
	t.Run("wait for interupt, shutdown gracefully", func(t *testing.T) {
		var (
			interrupt = make(chan os.Signal, 1)
			spyServer = newSpyServer()
			server    = gracefulshutdown.NewServer(spyServer, gracefulshutdown.WithSignal(interrupt))
		)

		spyServer.listenAndServeFn = func() error {
			return nil
		}

		spyServer.shutdownFn = func() error {
			return nil
		}

		go func() {
			if err := server.ListenAndServe(); err != nil {
				t.Error(err)
			}
		}()

		// verify we call listen on the delegate server
		spyServer.assertListen(t)

		// verify we call shutdown on the delegate server when an interrupt is made
		interrupt <- os.Interrupt
		spyServer.assertShutdown(t)
	})

	t.Run("when listen fails, return error", func(t *testing.T) {
		var (
			spyServer = newSpyServer()
			server    = gracefulshutdown.NewServer(spyServer)
			err       = errors.New("test error")
		)

		spyServer.listenAndServeFn = func() error {
			return err
		}

		gotErr := server.ListenAndServe()
		assert.Equal(t, errors.Is(gotErr, err), true)
	})

	t.Run("shutdown error gets propagated", func(t *testing.T) {
		var (
			interrupt = make(chan os.Signal)
			errChan   = make(chan error)
			spyServer = newSpyServer()
			server    = gracefulshutdown.NewServer(spyServer, gracefulshutdown.WithSignal(interrupt))
			err       = errors.New("test error")
		)

		spyServer.listenAndServeFn = func() error {
			return nil
		}

		spyServer.shutdownFn = func() error {
			return err
		}

		go func() {
			errChan <- server.ListenAndServe()
		}()

		interrupt <- os.Kill

		select {
		case gotErr := <-errChan:
			assert.Equal(t, errors.Is(gotErr, err), true)
		case <-time.After(500 * time.Millisecond):
			t.Error("timed out waiting for shutdown error to be propagated")
		}
	})
}

// implements: gracefulshutdown.HttpServer
type spyServer struct {
	listenAndServeFn func() error
	shutdownFn       func() error

	listened chan struct{}
	shutdown chan struct{}
}

func newSpyServer() *spyServer {
	return &spyServer{
		listened: make(chan struct{}, 1),
		shutdown: make(chan struct{}, 1),
	}
}

func (s *spyServer) ListenAndServe() error {
	s.listened <- struct{}{}
	return s.listenAndServeFn()
}

func (s *spyServer) Shutdown(ctx context.Context) error {
	s.shutdown <- struct{}{}
	return s.shutdownFn()
}

func (s *spyServer) assertListen(t testing.TB) {
	t.Helper()
	assert.SignalSent(t, s.listened, "listened")
}

func (s *spyServer) assertShutdown(t testing.TB) {
	t.Helper()
	assert.SignalSent(t, s.shutdown, "shutdown")
}
