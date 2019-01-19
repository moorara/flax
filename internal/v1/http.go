package v1

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"net/http"
	"path"
	"time"

	"github.com/gorilla/mux"
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

// SetDefaults set default values for empty fields
func (e *HTTPExpect) SetDefaults() {
	if e.Methods == nil || len(e.Methods) == 0 {
		e.Methods = []string{"GET"}
	}

	e.Path = path.Clean("/" + e.Path)

	if e.Queries == nil {
		e.Queries = map[string]string{}
	}

	if e.Headers == nil {
		e.Headers = map[string]string{}
	}
}

// Hash calculates a hash for an http expectation
func (e *HTTPExpect) Hash() uint64 {
	h := fnv.New64a()

	hashArray(h, true, e.Methods)
	hashString(h, e.Path)
	hashBool(h, e.Prefix)
	hashMap(h, true, e.Queries)
	hashMap(h, true, e.Headers)

	return h.Sum64()
}

// SetDefaults set default values for empty fields
func (r *HTTPResponse) SetDefaults() {
	if r.Delay == "" {
		r.Delay = "0"
	}

	if r.StatusCode < 100 || r.StatusCode > 599 {
		r.StatusCode = 200
	}

	if r.Headers == nil {
		r.Headers = map[string]string{}
	}
}

// GetHandler returns an http handler
func (r *HTTPResponse) GetHandler() http.HandlerFunc {
	d, _ := time.ParseDuration(r.Delay)

	return func(res http.ResponseWriter, req *http.Request) {
		time.Sleep(d)
		res.WriteHeader(r.StatusCode)
		for key, val := range r.Headers {
			res.Header().Add(key, val)
		}
		json.NewEncoder(res).Encode(r.Body)
	}
}

// SetDefaults set default values for empty fields
func (f *HTTPForward) SetDefaults() {
	if f.Delay == "" {
		f.Delay = "0"
	}

	if f.Headers == nil {
		f.Headers = map[string]string{}
	}
}

// GetHandler returns an http handler
func (f *HTTPForward) GetHandler() http.HandlerFunc {
	d, _ := time.ParseDuration(f.Delay)

	// TODO: implement proxy
	return func(res http.ResponseWriter, req *http.Request) {
		time.Sleep(d)
		res.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(res).Encode(JSON{
			"message": "this functionality is not yet available!",
		})
	}
}

// SetDefaults set default values for empty fields
func (m *HTTPMock) SetDefaults() {
	m.HTTPExpect.SetDefaults()

	if m.HTTPResponse == nil && m.HTTPForward == nil {
		m.HTTPResponse = &HTTPResponse{}
	}

	if m.HTTPResponse != nil {
		m.HTTPResponse.SetDefaults()
	}

	if m.HTTPForward != nil {
		m.HTTPForward.SetDefaults()
	}
}

// Hash calculates a hash for an http mock based on the http expectation
func (m *HTTPMock) Hash() uint64 {
	return m.HTTPExpect.Hash()
}

// RegisterRoutes configure routes for an http mock
func (m *HTTPMock) RegisterRoutes(router *mux.Router) {
	route := router.Methods(m.HTTPExpect.Methods...)

	if m.HTTPExpect.Prefix {
		route.PathPrefix(m.HTTPExpect.Path)
	} else {
		route.Path(m.HTTPExpect.Path)
	}

	for query, pattern := range m.HTTPExpect.Queries {
		route.Queries(query, fmt.Sprintf("{%s:%s}", query, pattern))
	}

	for header, pattern := range m.HTTPExpect.Headers {
		route.HeadersRegexp(header, pattern)
	}

	if m.HTTPResponse != nil {
		route.HandlerFunc(m.HTTPResponse.GetHandler())
	} else if m.HTTPForward != nil {
		route.HandlerFunc(m.HTTPForward.GetHandler())
	}
}
