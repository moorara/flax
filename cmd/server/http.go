package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/moorara/flax/cmd/config"
	"github.com/moorara/observe/log"
)

// HTTPMockServer is an http server for mocked http endpoints.
type HTTPMockServer struct {
	logger *log.Logger
	server *http.Server
}

// NewHTTPMockServer creates an http mock server.
func NewHTTPMockServer(logger *log.Logger, port string, handler http.Handler) *HTTPMockServer {
	return &HTTPMockServer{
		logger: logger,
		server: &http.Server{
			Addr:    port,
			Handler: handler,
		},
	}
}

// Start starts the server and blocks until either there is an error or the server is shutdown.
func (s *HTTPMockServer) Start() {
	done := make(chan struct{})

	go func() {
		// Catch OS signals
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigs
		s.logger.Infof("http mock server interrupted by signal %s", sig.String())

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

	s.logger.Infof("http mock server starting on %s ...", s.server.Addr)

	// ListenAndServe always returns a non-nil error.
	// After Shutdown or Close, the returned error is ErrServerClosed.
	err := s.server.ListenAndServe()
	if err != http.ErrServerClosed {
		s.logger.Errorf("http mock server errored: %s", err)
	}

	<-done
}
