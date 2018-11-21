package model

import (
	"hash/fnv"
	"net/http"
	"path"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/moorara/flax/pkg/log"
	"github.com/moorara/flax/pkg/metrics"
)

const (
	idPath = "/{id:[-0-9A-Za-z]+}"
)

type (
	// RESTExpect represents a RESTful expectation
	RESTExpect struct {
		BasePath string            `json:"basePath" yaml:"base_path"`
		Headers  map[string]string `json:"headers" yaml:"headers"`
	}

	// RESTResponse represents a mock RESTful response
	RESTResponse struct {
		Delay            string            `json:"delay" yaml:"delay"`
		PostStatusCode   int               `json:"postStatus" yaml:"post_status"`
		PutStatusCode    int               `json:"putStatus" yaml:"put_status"`
		PatchStatusCode  int               `json:"patchStatus" yaml:"patch_status"`
		DeleteStatusCode int               `json:"deleteStatus" yaml:"delete_status"`
		ListProperty     string            `json:"listProperty" yaml:"list_property"`
		Headers          map[string]string `json:"headers" yaml:"headers"`
	}

	// RESTStore represents a collection of RESTful resources
	RESTStore struct {
		Identifier string `json:"identifier" yaml:"identifier"`
		Objects    []JSON `json:"objects" yaml:"objects"`
	}

	// RESTMock represents a RESTful mock
	RESTMock struct {
		RESTExpect   `json:",inline" yaml:",inline"`
		RESTResponse `json:"response" yaml:"response"`
		RESTStore    `json:"store" yaml:"store"`
	}
)

// WithDefaults returns a rest expectation with default values
func (e RESTExpect) WithDefaults() RESTExpect {
	e.BasePath = "/" + e.BasePath
	e.BasePath = path.Clean(e.BasePath)

	if e.Headers == nil {
		e.Headers = map[string]string{}
	}

	return e
}

// Hash calculates a hash for a rest expectation based on the base path
func (e RESTExpect) Hash() uint64 {
	hash := fnv.New64a()
	hash.Write([]byte(e.BasePath))
	return hash.Sum64()
}

// WithDefaults returns a rest response with default values
func (r RESTResponse) WithDefaults() RESTResponse {
	if r.Delay == "" {
		r.Delay = "0"
	}

	if r.PostStatusCode < 100 || r.PostStatusCode > 599 {
		r.PostStatusCode = 201
	}

	if r.PutStatusCode < 100 || r.PutStatusCode > 599 {
		r.PutStatusCode = 200
	}

	if r.PatchStatusCode < 100 || r.PatchStatusCode > 599 {
		r.PatchStatusCode = 200
	}

	if r.DeleteStatusCode < 100 || r.DeleteStatusCode > 599 {
		r.DeleteStatusCode = 204
	}

	if r.ListProperty == "" {
		r.ListProperty = "" // returns a list of objects as an array
	}

	if r.Headers == nil {
		r.Headers = map[string]string{}
	}

	return r
}

// WithDefaults returns a rest store with default values
func (s RESTStore) WithDefaults() RESTStore {
	if s.Identifier == "" {
		s.Identifier = "" // will try from a standard list of identifiers
	}

	if s.Objects == nil {
		s.Objects = []JSON{}
	}

	return s
}

// WithDefaults returns a rest mock with default values
func (m RESTMock) WithDefaults() RESTMock {
	m.RESTExpect = m.RESTExpect.WithDefaults()
	m.RESTResponse = m.RESTResponse.WithDefaults()
	m.RESTStore = m.RESTStore.WithDefaults()

	return m
}

// Hash calculates a hash for a rest mock based on the rest expectation base path
func (m RESTMock) Hash() uint64 {
	return m.RESTExpect.Hash()
}

// RegisterRoute adds a new router to a Mux router for a rest mock
func (m RESTMock) RegisterRoute(router *mux.Router, logger *log.Logger, metrics *metrics.Metrics) {
	// GET /
	{
		path := m.RESTExpect.BasePath
		route := router.Methods("GET").Path(path)
		for header, pattern := range m.RESTExpect.Headers {
			route.HeadersRegexp(header, pattern)
		}

		route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		})
	}

	// POST /
	{
		path := m.RESTExpect.BasePath
		route := router.Methods("POST").Path(path)
		for header, pattern := range m.RESTExpect.Headers {
			route.HeadersRegexp(header, pattern)
		}

		route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		})
	}

	// GET /
	{
		path := filepath.Join(m.RESTExpect.BasePath, idPath)
		route := router.Methods("GET").Path(path)
		for header, pattern := range m.RESTExpect.Headers {
			route.HeadersRegexp(header, pattern)
		}

		route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		})
	}

	// PUT /
	{
		path := filepath.Join(m.RESTExpect.BasePath, idPath)
		route := router.Methods("PUT").Path(path)
		for header, pattern := range m.RESTExpect.Headers {
			route.HeadersRegexp(header, pattern)
		}

		route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		})
	}

	// PATCH /
	{
		path := filepath.Join(m.RESTExpect.BasePath, idPath)
		route := router.Methods("PATCH").Path(path)
		for header, pattern := range m.RESTExpect.Headers {
			route.HeadersRegexp(header, pattern)
		}

		route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		})
	}

	// DELETE /
	{
		path := filepath.Join(m.RESTExpect.BasePath, idPath)
		route := router.Methods("DELETE").Path(path)
		for header, pattern := range m.RESTExpect.Headers {
			route.HeadersRegexp(header, pattern)
		}

		route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		})
	}
}
