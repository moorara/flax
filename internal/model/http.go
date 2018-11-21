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
		Prefix  bool              `json:"prefix" yaml:"prefix"`
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
		HTTPExpect    `json:",inline" yaml:",inline"`
		*HTTPResponse `json:"response" yaml:"response"`
		*HTTPForward  `json:"forward" yaml:"forward"`
	}
)

// WithDefaults returns an http expectation with default values
func (e HTTPExpect) WithDefaults() HTTPExpect {
	if e.Methods == nil || len(e.Methods) == 0 {
		e.Methods = []string{"GET"}
	}

	e.Path = path.Clean("/" + e.Path)

	// Prefix is a boolean defaulting to false

	if e.Queries == nil {
		e.Queries = map[string]string{}
	}

	if e.Headers == nil {
		e.Headers = map[string]string{}
	}

	return e
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

	// Body

	return r
}

// WithDefaults returns an http forward with default values
func (f HTTPForward) WithDefaults() HTTPForward {
	if f.Delay == "" {
		f.Delay = "0"
	}

	// To

	if f.Headers == nil {
		f.Headers = map[string]string{}
	}

	return f
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

// Hash calculates a hash for an http mock based on the http expectation values
func (m HTTPMock) Hash() uint64 {
	hash := fnv.New64a()

	for _, m := range m.HTTPExpect.Methods {
		hash.Write([]byte(m))
	}

	hash.Write([]byte(m.HTTPExpect.Path))

	queries := []Pair{}
	for name, value := range m.HTTPExpect.Queries {
		queries = append(queries, Pair{
			Name:  name,
			Value: value,
		})
	}

	for _, q := range queries {
		hash.Write([]byte(q.Name))
		hash.Write([]byte(q.Value))
	}

	headers := []Pair{}
	for name, value := range m.HTTPExpect.Headers {
		headers = append(headers, Pair{
			Name:  name,
			Value: value,
		})
	}

	for _, h := range headers {
		hash.Write([]byte(h.Name))
		hash.Write([]byte(h.Value))
	}

	return hash.Sum64()
}
