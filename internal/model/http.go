package model

import (
	"hash/fnv"
	"path"
)

type (
	// HTTPExpect represents an http expectation
	HTTPExpect struct {
		Methods []string          `json:"methods" yaml:"methods"`
		Path    string            `json:"path" yaml:"path"`
		Queries map[string]string `json:"queries" yaml:"queries"`
		Headers map[string]string `json:"headers" yaml:"headers"`
	}

	// HTTPResponse represents a mock http response
	HTTPResponse struct {
		Delay      string            `json:"delay" yaml:"delay"`
		StatusCode int               `json:"status" yaml:"status"`
		Headers    map[string]string `json:"headers" yaml:"headers"`
		Body       interface{}       `json:"body" yaml:"body"`
	}

	// HTTPForward represents a forwarder for an http request
	HTTPForward struct {
		Delay   string            `json:"delay" yaml:"delay"`
		To      string            `json:"to" yaml:"to"`
		Headers map[string]string `json:"headers" yaml:"headers"`
	}

	// HTTPMock represents an http mock
	HTTPMock struct {
		HTTPExpect
		*HTTPResponse `json:"response" yaml:"response"`
		*HTTPForward  `json:"forward" yaml:"forward"`
	}
)

// WithDefaults returns an http expectation with default values
func (e HTTPExpect) WithDefaults() HTTPExpect {
	if e.Methods == nil || len(e.Methods) == 0 {
		e.Methods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	}

	e.Path = "/" + e.Path
	e.Path = path.Clean(e.Path)

	if e.Queries == nil {
		e.Queries = map[string]string{}
	}

	if e.Headers == nil {
		e.Headers = map[string]string{}
	}

	return e
}

// Hash calculates a hash for an http expectation
func (e HTTPExpect) Hash() uint64 {
	hash := fnv.New64a()

	return hash.Sum64()
}

// WithDefaults returns an http response with default values
func (r HTTPResponse) WithDefaults() HTTPResponse {
	if r.Delay == "" {
		r.Delay = "0"
	}

	if r.StatusCode < 100 || r.StatusCode > 599 {
		r.StatusCode = 200
	}

	if r.Headers == nil {
		r.Headers = map[string]string{}
	}

	return r
}

// Hash calculates a hash for an http response
func (r HTTPResponse) Hash() uint64 {
	hash := fnv.New64a()

	return hash.Sum64()
}

// WithDefaults returns an http forward with default values
func (f HTTPForward) WithDefaults() HTTPForward {
	if f.Delay == "" {
		f.Delay = "0"
	}

	if f.Headers == nil {
		f.Headers = map[string]string{}
	}

	return f
}

// Hash calculates a hash for an http forward
func (f HTTPForward) Hash() uint64 {
	hash := fnv.New64a()

	return hash.Sum64()
}

// WithDefaults returns an http mock with default values
func (m HTTPMock) WithDefaults() HTTPMock {
	m.HTTPExpect = m.HTTPExpect.WithDefaults()

	if m.HTTPResponse != nil {
		hr := m.HTTPResponse.WithDefaults()
		m.HTTPResponse = &hr
	}

	if m.HTTPForward != nil {
		hf := m.HTTPForward.WithDefaults()
		m.HTTPForward = &hf
	}

	return m
}

// Hash calculates a hash for an http mock
func (m HTTPMock) Hash() uint64 {
	hash := m.HTTPExpect.Hash()

	if m.HTTPResponse != nil {
		hash += m.HTTPResponse.Hash()
	}

	if m.HTTPForward != nil {
		hash += m.HTTPForward.Hash()
	}

	return hash
}
