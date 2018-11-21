package model

import (
	"github.com/gorilla/mux"
	"github.com/moorara/flax/pkg/log"
	"github.com/moorara/flax/pkg/metrics"
)

type (
	// JSON is the type for json objects
	JSON map[string]interface{}

	// Mock is the interface for a mock
	Mock interface {
		Hash() uint64
		RegisterRoute(*mux.Router, *log.Logger, *metrics.Metrics)
	}

	// Pair is a name-value pair
	Pair struct {
		Name  string `json:"name" yaml:"name"`
		Value string `json:"value" yaml:"value"`
	}
)
