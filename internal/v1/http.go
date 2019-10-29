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

// HTTPExpect represents an http expectation.
type HTTPExpect struct {
	Methods []string          `json:"methods" yaml:"methods"`
	Path    string            `json:"path" yaml:"path"`
	Prefix  bool              `json:"prefix" yaml:"prefix"`
	Queries map[string]string `json:"queries" yaml:"queries"`
	Headers map[string]string `json:"headers" yaml:"headers"`
}

// Hash calculates a hash for an http expectation.
func (e HTTPExpect) Hash() uint64 {
	h := fnv.New64a()

	hashStringSlice(h, true, e.Methods)
	hashString(h, e.Path)
	hashBool(h, e.Prefix)
	hashStringMap(h, true, e.Queries)
	hashStringMap(h, true, e.Headers)

	return h.Sum64()
}

// SetDefaults set default values for empty fields.
func (e *HTTPExpect) SetDefaults() {
	if len(e.Methods) == 0 {
		e.Methods = []string{"GET"}
	}

	e.Path = path.Clean("/" + e.Path)
}

// HTTPResponse represents a mock http response.
type HTTPResponse struct {
	Delay      string            `json:"delay" yaml:"delay"`
	StatusCode int               `json:"status" yaml:"status"`
	Headers    map[string]string `json:"headers" yaml:"headers"`
	Body       interface{}       `json:"body" yaml:"body"`
}

// SetDefaults set default values for empty fields.
func (r *HTTPResponse) SetDefaults() {
	if r.StatusCode < 100 || r.StatusCode > 599 {
		r.StatusCode = 200
	}
}

// Handler returns an http handler.
func (r HTTPResponse) Handler() http.HandlerFunc {
	d, _ := time.ParseDuration(r.Delay)

	return func(res http.ResponseWriter, req *http.Request) {
		time.Sleep(d)
		res.WriteHeader(r.StatusCode)
		for key, val := range r.Headers {
			res.Header().Set(key, val)
		}
		json.NewEncoder(res).Encode(r.Body)
	}
}

// HTTPForward represents a forwarder for an http request.
type HTTPForward struct {
	Delay   string            `json:"delay" yaml:"delay"`
	To      string            `json:"to" yaml:"to"`
	Headers map[string]string `json:"headers" yaml:"headers"`
}

// SetDefaults set default values for empty fields.
func (f *HTTPForward) SetDefaults() {
	// Nothing to set as default
}

// Handler returns an http handler.
func (f HTTPForward) Handler() http.HandlerFunc {
	d, _ := time.ParseDuration(f.Delay)

	// TODO: implement proxy
	return func(res http.ResponseWriter, req *http.Request) {
		time.Sleep(d)
		res.WriteHeader(http.StatusNotImplemented)
		for key, val := range f.Headers {
			res.Header().Add(key, val)
		}

		json.NewEncoder(res).Encode(JSON{
			"message": "this functionality is not yet available!",
		})
	}
}

// HTTPMock represents an http mock.
type HTTPMock struct {
	HTTPExpect    `json:",inline" yaml:",inline"`
	*HTTPResponse `json:"response" yaml:"response"`
	*HTTPForward  `json:"forward" yaml:"forward"`
}

// Hash calculates a hash for an http mock based on the http expectation.
func (m HTTPMock) Hash() uint64 {
	return m.HTTPExpect.Hash()
}

// SetDefaults set default values for empty fields.
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

// RegisterRoutes configure routes for an http mock.
func (m HTTPMock) RegisterRoutes(router *mux.Router) {
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
		route.HandlerFunc(m.HTTPResponse.Handler())
	} else if m.HTTPForward != nil {
		route.HandlerFunc(m.HTTPForward.Handler())
	}
}

// DefaultHTTPMock returns a default HTTPMock.
func DefaultHTTPMock() HTTPMock {
	return HTTPMock{
		HTTPExpect: HTTPExpect{
			Methods: []string{"GET"},
			Path:    "/",
		},
		HTTPResponse: &HTTPResponse{
			StatusCode: 200,
		},
	}
}
