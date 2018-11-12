package model

import (
	"net/http"
)

type (
	// HTTPFilters is the model for filtering http requests
	HTTPFilters struct {
		Params  map[string]string `json:"params" yaml:"params"`
		Headers map[string]string `json:"headers" yaml:"headers"`
	}

	// HTTPExpectation is the model for an http expectation
	HTTPExpectation struct {
		Methods []string    `json:"methods" yaml:"methods"`
		Path    string      `json:"path" yaml:"path"`
		Filters HTTPFilters `json:"filters" yaml:"filters"`
	}

	// HTTPEndpoint is the model an http endpoint
	HTTPEndpoint struct {
		HTTPExpectation
		Handler http.HandlerFunc
	}
)
