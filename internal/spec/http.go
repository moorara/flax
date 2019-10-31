package spec

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"net/http"
	"path"
	"strings"
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

// HTTPResponse represents a mock http response.
type HTTPResponse struct {
	Delay      string            `json:"delay" yaml:"delay"`
	StatusCode int               `json:"status" yaml:"status"`
	Headers    map[string]string `json:"headers" yaml:"headers"`
	Body       interface{}       `json:"body" yaml:"body"`
}

// HTTPForward represents a forwarder for an http request.
type HTTPForward struct {
	Delay   string            `json:"delay" yaml:"delay"`
	To      string            `json:"to" yaml:"to"`
	Headers map[string]string `json:"headers" yaml:"headers"`
}

// HTTPMock represents an http mock.
type HTTPMock struct {
	HTTPExpect    `json:",inline" yaml:",inline"`
	*HTTPResponse `json:"response" yaml:"response"`
	*HTTPForward  `json:"forward" yaml:"forward"`
}

// SetDefaults set default values for empty fields.
func (m *HTTPMock) SetDefaults() {
	if len(m.HTTPExpect.Methods) == 0 {
		m.HTTPExpect.Methods = []string{"GET"}
	}

	m.HTTPExpect.Path = path.Clean("/" + m.HTTPExpect.Path)

	if m.HTTPResponse == nil && m.HTTPForward == nil {
		m.HTTPResponse = &HTTPResponse{}
	}

	if m.HTTPResponse != nil {
		if m.HTTPResponse.StatusCode == 0 {
			m.HTTPResponse.StatusCode = 200
		}
	}

	if m.HTTPForward != nil {
		// No default
	}
}

// String returns a string representation of the mock.
func (m HTTPMock) String() string {
	return fmt.Sprintf(
		"%s %s",
		strings.Join(m.HTTPExpect.Methods, "|"),
		m.HTTPExpect.Path,
	)
}

// Hash calculates a hash for an http mock based on the http expectation.
func (m HTTPMock) Hash() uint64 {
	h := fnv.New64a()

	hashStringSlice(h, true, m.HTTPExpect.Methods)
	hashString(h, m.HTTPExpect.Path)
	hashBool(h, m.HTTPExpect.Prefix)
	hashStringMap(h, true, m.HTTPExpect.Queries)
	hashStringMap(h, true, m.HTTPExpect.Headers)

	return h.Sum64()
}

// RegisterRoutes configure routes for an http mock.
func (m HTTPMock) RegisterRoutes(router *mux.Router) {
	route := router.NewRoute()

	route.Methods(m.HTTPExpect.Methods...)

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
		responseDelay, _ := time.ParseDuration(m.HTTPResponse.Delay)
		route.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			time.Sleep(responseDelay)
			res.WriteHeader(m.HTTPResponse.StatusCode)
			for key, val := range m.HTTPResponse.Headers {
				res.Header().Set(key, val)
			}
			_ = json.NewEncoder(res).Encode(m.HTTPResponse.Body)
		})
	} else if m.HTTPForward != nil {
		forwardDelay, _ := time.ParseDuration(m.HTTPForward.Delay)
		route.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			// TODO: implement proxy
			time.Sleep(forwardDelay)
			res.WriteHeader(http.StatusNotImplemented)
			_ = json.NewEncoder(res).Encode(JSON{
				"message": "this functionality is not yet available!",
			})
		})
	}
}
