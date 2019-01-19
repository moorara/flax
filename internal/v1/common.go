package v1

import (
	"github.com/gorilla/mux"
)

type (
	// JSON is the type for json objects
	JSON map[string]interface{}

	// Mock is the interface for a mock object
	Mock interface {
		Hash() uint64
		RegisterRoutes(*mux.Router)
	}

	// Pair is a name-value pair
	Pair struct {
		Name  string `json:"name" yaml:"name"`
		Value string `json:"value" yaml:"value"`
	}
)
