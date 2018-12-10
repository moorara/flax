package server

import (
	"context"
)

type (
	// HTTPServer is the interface for http.Server
	HTTPServer interface {
		ListenAndServe() error
		Shutdown(context.Context) error
	}

	// Server is the generic type for a server
	Server interface {
		Start()
		IsLive() bool
		IsReady() bool
	}
)
