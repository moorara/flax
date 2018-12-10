package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/moorara/flax/pkg/log"
)

const (
	// Default Kubernetes grace period
	timeout = 30 * time.Second
)

// ControlServer is an http server for configuring Flax
type ControlServer struct {
	logger     *log.Logger
	httpServer HTTPServer
}

// NewControlServer creates a control server
func NewControlServer(logger *log.Logger, port string, router *mux.Router) *ControlServer {
	return &ControlServer{
		logger: logger,
		httpServer: &http.Server{
			Addr:    port,
			Handler: router,
		},
	}
}

// Start starts the server
func (s *ControlServer) Start() {
	done := make(chan struct{})

	go func() {
		// Catch OS signals
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigs
		s.logger.Error("message", fmt.Sprintf("control server interrupted by signal %s", sig.String()))

		// Shutdown http server gracefully
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		err := s.httpServer.Shutdown(ctx)
		if err != nil {
			s.logger.Error("message", "control server failed to gracefully shutdown.", "error", err)
		}
		s.logger.Info("message", "control server gracefully shutdown.")

		close(done)
	}()

	// Start http server
	s.logger.Info("message", "control server started ...")
	err := s.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.logger.Error("message", "control server errored.", "error", err)
	}

	<-done
}

// IsLive determines if the server is live
func (s *ControlServer) IsLive() bool {
	return true
}

// IsReady determines if the server is ready
func (s *ControlServer) IsReady() bool {
	return true
}
