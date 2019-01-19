package v1

import (
	"hash/fnv"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
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
		GetStatusCode    int               `json:"getStatus" yaml:"get_status"`
		PostStatusCode   int               `json:"postStatus" yaml:"post_status"`
		PutStatusCode    int               `json:"putStatus" yaml:"put_status"`
		PatchStatusCode  int               `json:"patchStatus" yaml:"patch_status"`
		DeleteStatusCode int               `json:"deleteStatus" yaml:"delete_status"`
		ListProperty     string            `json:"listProperty" yaml:"list_property"`
		Headers          map[string]string `json:"headers" yaml:"headers"`
	}

	// RESTStore represents a collection of RESTful resources
	RESTStore struct {
		Identifier string               `json:"identifier" yaml:"identifier"`
		Objects    []JSON               `json:"objects" yaml:"objects"`
		Directory  map[interface{}]JSON `json:"-" yaml:"-"`
	}

	// RESTMock represents a RESTful mock
	RESTMock struct {
		RESTExpect   `json:",inline" yaml:",inline"`
		RESTResponse `json:"response" yaml:"response"`
		RESTStore    `json:"store" yaml:"store"`
	}
)

// SetDefaults set default values for empty fields
func (e *RESTExpect) SetDefaults() {
	e.BasePath = path.Clean("/" + e.BasePath)

	if e.Headers == nil {
		e.Headers = map[string]string{}
	}
}

// Hash calculates a hash for a rest expectation
func (e *RESTExpect) Hash() uint64 {
	h := fnv.New64a()

	hashString(h, e.BasePath)
	hashMap(h, true, e.Headers)

	return h.Sum64()
}

// SetDefaults set default values for empty fields
func (r *RESTResponse) SetDefaults() {
	if r.Delay == "" {
		r.Delay = "0"
	}

	if r.GetStatusCode < 100 || r.GetStatusCode > 599 {
		r.GetStatusCode = 200
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
		r.ListProperty = ""
	}

	if r.Headers == nil {
		r.Headers = map[string]string{}
	}
}

// SetDefaults set default values for empty fields
func (s *RESTStore) SetDefaults() {
	if s.Identifier == "" {
		s.Identifier = ""
	}

	if s.Objects == nil {
		s.Objects = []JSON{}
	}

	s.Directory = map[interface{}]JSON{}
	for _, obj := range s.Objects {
		val, err := findID(s.Identifier, obj)
		if err == nil {
			s.Directory[val] = obj
		}
	}
}

// SetDefaults set default values for empty fields
func (m *RESTMock) SetDefaults() {
	m.RESTExpect.SetDefaults()
	m.RESTResponse.SetDefaults()
	m.RESTStore.SetDefaults()
}

// Hash calculates a hash for a rest mock based on the rest expectation
func (m *RESTMock) Hash() uint64 {
	return m.RESTExpect.Hash()
}

// RegisterRoutes configure routes for a rest mock
func (m *RESTMock) RegisterRoutes(router *mux.Router) {
	d, _ := time.ParseDuration(m.Delay)

	// GET /
	{
		path := m.RESTExpect.BasePath
		route := router.Methods("GET").Path(path)
		for header, pattern := range m.RESTExpect.Headers {
			route.HeadersRegexp(header, pattern)
		}

		// TODO:
		route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if m.RESTStore.Directory != nil {
				log.Printf("Getting all ...\n")
				time.Sleep(d)
			}
		})
	}

	// POST /
	{
		path := m.RESTExpect.BasePath
		route := router.Methods("POST").Path(path)
		for header, pattern := range m.RESTExpect.Headers {
			route.HeadersRegexp(header, pattern)
		}

		// TODO:
		route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if m.RESTStore.Directory != nil {
				log.Printf("Creating new ...\n")
				time.Sleep(d)
			}
		})
	}

	// GET /
	{
		path := filepath.Join(m.RESTExpect.BasePath, idPath)
		route := router.Methods("GET").Path(path)
		for header, pattern := range m.RESTExpect.Headers {
			route.HeadersRegexp(header, pattern)
		}

		// TODO:
		route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			id := vars["id"]

			if m.RESTStore.Directory != nil {
				log.Printf("Getting %s ...\n", id)
				time.Sleep(d)
			}
		})
	}

	// PUT /
	{
		path := filepath.Join(m.RESTExpect.BasePath, idPath)
		route := router.Methods("PUT").Path(path)
		for header, pattern := range m.RESTExpect.Headers {
			route.HeadersRegexp(header, pattern)
		}

		// TODO:
		route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			id := vars["id"]

			if m.RESTStore.Directory != nil {
				log.Printf("Deleting %s ...\n", id)
				time.Sleep(d)
			}
		})
	}

	// PATCH /
	{
		path := filepath.Join(m.RESTExpect.BasePath, idPath)
		route := router.Methods("PATCH").Path(path)
		for header, pattern := range m.RESTExpect.Headers {
			route.HeadersRegexp(header, pattern)
		}

		// TODO:
		route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			id := vars["id"]

			if m.RESTStore.Directory != nil {
				log.Printf("Patching %s ...\n", id)
				time.Sleep(d)
			}
		})
	}

	// DELETE /
	{
		path := filepath.Join(m.RESTExpect.BasePath, idPath)
		route := router.Methods("DELETE").Path(path)
		for header, pattern := range m.RESTExpect.Headers {
			route.HeadersRegexp(header, pattern)
		}

		// TODO:
		route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			id := vars["id"]

			if m.RESTStore.Directory != nil {
				log.Printf("Deleting %s ...\n", id)
				time.Sleep(d)
			}
		})
	}
}
