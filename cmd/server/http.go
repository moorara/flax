package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/moorara/flax/pkg/log"
)

// HTTPMockServer is an http server for mocked http endpoints
type HTTPMockServer struct {
	logger     *log.Logger
	httpServer HTTPServer
}

// NewHTTPMockServer creates an http mock server
func NewHTTPMockServer(logger *log.Logger, port string, router *mux.Router) *HTTPMockServer {
	return &HTTPMockServer{
		logger: logger,
		httpServer: &http.Server{
			Addr:    port,
			Handler: router,
		},
	}
}

// Start starts the server
func (s *HTTPMockServer) Start() {
	done := make(chan struct{})

	go func() {
		// Catch OS signals
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigs
		s.logger.Error("message", fmt.Sprintf("http mock server interrupted by signal %s", sig.String()))

		// Shutdown http server gracefully
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		if err := s.httpServer.Shutdown(ctx); err != nil {
			s.logger.Error("message", "http mock server failed to gracefully shutdown.", "error", err)
		}
		s.logger.Info("message", "http mock server gracefully shutdown.")

		close(done)
	}()

	// Start http server
	s.logger.Info("message", "http mock server started ...")
	err := s.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.logger.Error("message", "http mock server errored.", "error", err)
	}

	<-done
}

// IsLive determines if the server is live
func (s *HTTPMockServer) IsLive() bool {
	return true
}

// IsReady determines if the server is ready
func (s *HTTPMockServer) IsReady() bool {
	return true
}
