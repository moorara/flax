package server

import (
	"context"
	"net/http"
	"testing"

	"github.com/moorara/log"
	"github.com/stretchr/testify/assert"
)

type mockHTTPServer struct {
	ListenAndServeOutError error
	ShutdownInContext      context.Context
	ShutdownOutError       error
}

func (m *mockHTTPServer) ListenAndServe() error {
	return m.ListenAndServeOutError
}

func (m *mockHTTPServer) Shutdown(ctx context.Context) error {
	m.ShutdownInContext = ctx
	return m.ShutdownOutError
}

func TestNewAPIServer(t *testing.T) {
	tests := []struct {
		name    string
		logger  log.Logger
		port    uint16
		handler http.Handler
	}{
		{
			"OK",
			log.NewNopLogger(),
			8080,
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			apiServer := NewAPIServer(tc.logger, tc.port, tc.handler)

			assert.NotNil(t, apiServer)
			assert.NotNil(t, apiServer.server)
			assert.Equal(t, tc.logger, apiServer.logger)
		})
	}
}

func TestAPIServerStart(t *testing.T) {
	tests := []struct {
		name          string
		server        *mockHTTPServer
		expectedError error
	}{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			apiServer := &APIServer{
				logger: log.NewNopLogger(),
				server: tc.server,
			}

			apiServer.Start()
		})
	}
}
