package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/moorara/flax/cmd/config"
	"github.com/moorara/log"
)

// HTTPServer is the interface for http.Server
type HTTPServer interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

// APIServer is an http server for mocked http endpoints.
type APIServer struct {
	logger log.Logger
	server HTTPServer
}

// NewAPIServer creates an http mock server.
func NewAPIServer(logger log.Logger, port uint16, handler http.Handler) *APIServer {
	addr := fmt.Sprintf(":%d", port)

	return &APIServer{
		logger: logger,
		server: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

// Start starts the server and blocks until either there is an error or the server is shutdown.
func (s *APIServer) Start() {
	done := make(chan struct{})

	go func() {
		// Catch OS signals
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigs
		s.logger.Infof("http mock server received signal %s", sig.String())

		ctx, cancel := context.WithTimeout(context.Background(), config.Global.GracePeriod)
		defer cancel()

		err := s.server.Shutdown(ctx)
		if err != nil {
			s.logger.Errorf("http mock server failed to gracefully shutdown: %s", err)
		} else {
			s.logger.Info("http mock server was gracefully shutdown.")
		}

		close(done)
	}()

	s.logger.Info("http mock server starting ...")

	// ListenAndServe always returns a non-nil error.
	// After Shutdown or Close, the returned error is ErrServerClosed.
	err := s.server.ListenAndServe()
	if err != http.ErrServerClosed {
		s.logger.Errorf("http mock server errored: %s", err)
	}

	<-done
}
