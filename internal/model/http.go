package model

import (
	"fmt"
	"hash/fnv"
	"net/http"
	"path"
	"sort"

	"github.com/gorilla/mux"
	"github.com/moorara/flax/pkg/log"
	"github.com/moorara/flax/pkg/metrics"
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
		HTTPExpect   `json:",inline" yaml:",inline"`
		HTTPResponse `json:"response" yaml:"response"`
		HTTPForward  `json:"forward" yaml:"forward"`
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

// Hash calculates a hash for an http expectation
func (e HTTPExpect) Hash() uint64 {
	hash := fnv.New64a()

	sort.Strings(e.Methods)
	for _, m := range e.Methods {
		hash.Write([]byte(m))
	}

	hash.Write([]byte(e.Path))

	if e.Prefix {
		hash.Write([]byte("true"))
	} else {
		hash.Write([]byte("false"))
	}

	queries := []Pair{}
	for name, value := range e.Queries {
		queries = append(queries, Pair{
			Name:  name,
			Value: value,
		})
	}

	sort.Slice(queries, func(i, j int) bool {
		return queries[i].Name < queries[j].Name
	})

	for _, q := range queries {
		hash.Write([]byte(q.Name))
		hash.Write([]byte(q.Value))
	}

	headers := []Pair{}
	for name, value := range e.Headers {
		headers = append(headers, Pair{
			Name:  name,
			Value: value,
		})
	}

	sort.Slice(headers, func(i, j int) bool {
		return headers[i].Name < headers[j].Name
	})

	for _, h := range headers {
		hash.Write([]byte(h.Name))
		hash.Write([]byte(h.Value))
	}

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
	m.HTTPResponse = m.HTTPResponse.WithDefaults()
	m.HTTPForward = m.HTTPForward.WithDefaults()

	return m
}

// Hash calculates a hash for an http mock based on the http expectation values
func (m HTTPMock) Hash() uint64 {
	return m.HTTPExpect.Hash()
}

// RegisterRoute adds a new router to a Mux router for an http mock
func (m HTTPMock) RegisterRoute(router *mux.Router, logger *log.Logger, metrics *metrics.Metrics) {
	route := router.Methods(m.HTTPExpect.Methods...)

	if m.HTTPExpect.Prefix {
		route.Path(m.HTTPExpect.Path)
	} else {
		route.PathPrefix(m.HTTPExpect.Path)
	}

	for query, pattern := range m.HTTPExpect.Queries {
		route.Queries(query, fmt.Sprintf("{%s:%s}", query, pattern))
	}

	for header, pattern := range m.HTTPExpect.Headers {
		route.HeadersRegexp(header, pattern)
	}

	route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
