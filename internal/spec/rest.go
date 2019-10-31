package spec

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"net/http"
	"path"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

const idTemplate = "/{id:[-0-9A-Za-z]+}"

// RESTExpect represents a RESTful expectation.
type RESTExpect struct {
	BasePath string            `json:"basePath" yaml:"base_path"`
	Headers  map[string]string `json:"headers" yaml:"headers"`
}

// RESTResponse represents a mock RESTful response.
type RESTResponse struct {
	Delay            string            `json:"delay" yaml:"delay"`
	GetStatusCode    int               `json:"getStatus" yaml:"get_status"`
	PostStatusCode   int               `json:"postStatus" yaml:"post_status"`
	PutStatusCode    int               `json:"putStatus" yaml:"put_status"`
	PatchStatusCode  int               `json:"patchStatus" yaml:"patch_status"`
	DeleteStatusCode int               `json:"deleteStatus" yaml:"delete_status"`
	Headers          map[string]string `json:"headers" yaml:"headers"`
	ListKey          string            `json:"listKey" yaml:"list_key"`
}

// RESTStore represents a collection of RESTful resources.
type RESTStore struct {
	Identifier string               `json:"identifier" yaml:"identifier"`
	Objects    []JSON               `json:"objects" yaml:"objects"`
	Directory  map[interface{}]JSON `json:"-" yaml:"-"`
}

// Index creates a map of identifiers to objects.
func (s *RESTStore) Index() {
	s.Directory = map[interface{}]JSON{}

	for _, obj := range s.Objects {
		val, err := findID(s.Identifier, obj)
		if err == nil {
			s.Directory[val] = obj
		}
	}
}

// RESTMock represents a RESTful mock.
type RESTMock struct {
	RESTExpect   `json:",inline" yaml:",inline"`
	RESTResponse `json:"response" yaml:"response"`
	RESTStore    `json:"store" yaml:"store"`
}

// SetDefaults set default values for empty fields.
func (m *RESTMock) SetDefaults() {
	m.RESTExpect.BasePath = path.Clean("/" + m.RESTExpect.BasePath)

	if m.RESTResponse.GetStatusCode == 0 {
		m.RESTResponse.GetStatusCode = 200
	}

	if m.RESTResponse.PostStatusCode == 0 {
		m.RESTResponse.PostStatusCode = 201
	}

	if m.RESTResponse.PutStatusCode == 0 {
		m.RESTResponse.PutStatusCode = 200
	}

	if m.RESTResponse.PatchStatusCode == 0 {
		m.RESTResponse.PatchStatusCode = 200
	}

	if m.RESTResponse.DeleteStatusCode == 0 {
		m.RESTResponse.DeleteStatusCode = 204
	}

	if m.RESTResponse.Headers == nil {
		m.RESTResponse.Headers = map[string]string{
			"Content-Type": "application/json",
		}
	}

	if m.RESTStore.Objects == nil {
		m.RESTStore.Objects = []JSON{}
	}
}

// String returns a string representation of the mock.
func (m RESTMock) String() string {
	return fmt.Sprintf("%s", m.RESTExpect.BasePath)
}

// Hash calculates a hash for a rest mock based on the rest expectation.
func (m RESTMock) Hash() uint64 {
	h := fnv.New64a()

	hashString(h, m.RESTExpect.BasePath)
	hashStringMap(h, true, m.RESTExpect.Headers)

	return h.Sum64()
}

// RegisterRoutes configure routes for a rest mock.
func (m RESTMock) RegisterRoutes(router *mux.Router) {
	delay, _ := time.ParseDuration(m.Delay)

	// GET /
	{
		path := m.RESTExpect.BasePath
		route := router.Methods("GET").Path(path)
		for header, pattern := range m.RESTExpect.Headers {
			route.HeadersRegexp(header, pattern)
		}

		// TODO: implement filtering through query parameters
		route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(m.RESTResponse.GetStatusCode)
			for key, val := range m.RESTResponse.Headers {
				w.Header().Set(key, val)
			}

			var resp interface{}
			if m.RESTResponse.ListKey == "" {
				resp = m.RESTStore.Objects
			} else {
				resp = JSON{
					m.RESTResponse.ListKey: m.RESTStore.Objects,
				}
			}

			_ = json.NewEncoder(w).Encode(resp)
		})
	}

	// POST /
	{
		path := m.RESTExpect.BasePath
		route := router.Methods("POST").Path(path)
		for header, pattern := range m.RESTExpect.Headers {
			route.HeadersRegexp(header, pattern)
		}

		// TODO: Finish implementation
		route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if m.RESTStore.Directory != nil {
				time.Sleep(delay)
				w.WriteHeader(m.RESTResponse.PostStatusCode)
				for key, val := range m.RESTResponse.Headers {
					w.Header().Set(key, val)
				}

				_ = json.NewEncoder(w).Encode(JSON{
					"message": "not implemented yet!",
				})
			}
		})
	}

	// GET /id
	{
		path := filepath.Join(m.RESTExpect.BasePath, idTemplate)
		route := router.Methods("GET").Path(path)
		for header, pattern := range m.RESTExpect.Headers {
			route.HeadersRegexp(header, pattern)
		}

		// TODO: Finish implementation
		route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// vars := mux.Vars(r)
			// id := vars["id"]

			if m.RESTStore.Directory != nil {
				time.Sleep(delay)
				w.WriteHeader(m.RESTResponse.GetStatusCode)
				for key, val := range m.RESTResponse.Headers {
					w.Header().Set(key, val)
				}

				_ = json.NewEncoder(w).Encode(JSON{
					"message": "not implemented yet!",
				})
			}
		})
	}

	// PUT /id
	{
		path := filepath.Join(m.RESTExpect.BasePath, idTemplate)
		route := router.Methods("PUT").Path(path)
		for header, pattern := range m.RESTExpect.Headers {
			route.HeadersRegexp(header, pattern)
		}

		// TODO: Finish implementation
		route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// vars := mux.Vars(r)
			// id := vars["id"]

			if m.RESTStore.Directory != nil {
				time.Sleep(delay)
				w.WriteHeader(m.RESTResponse.PutStatusCode)
				for key, val := range m.RESTResponse.Headers {
					w.Header().Set(key, val)
				}

				_ = json.NewEncoder(w).Encode(JSON{
					"message": "not implemented yet!",
				})
			}
		})
	}

	// PATCH /id
	{
		path := filepath.Join(m.RESTExpect.BasePath, idTemplate)
		route := router.Methods("PATCH").Path(path)
		for header, pattern := range m.RESTExpect.Headers {
			route.HeadersRegexp(header, pattern)
		}

		// TODO: Finish implementation
		route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// vars := mux.Vars(r)
			// id := vars["id"]

			if m.RESTStore.Directory != nil {
				time.Sleep(delay)
				w.WriteHeader(m.RESTResponse.PatchStatusCode)
				for key, val := range m.RESTResponse.Headers {
					w.Header().Set(key, val)
				}

				_ = json.NewEncoder(w).Encode(JSON{
					"message": "not implemented yet!",
				})
			}
		})
	}

	// DELETE /id
	{
		path := filepath.Join(m.RESTExpect.BasePath, idTemplate)
		route := router.Methods("DELETE").Path(path)
		for header, pattern := range m.RESTExpect.Headers {
			route.HeadersRegexp(header, pattern)
		}

		// TODO: Finish implementation
		route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// vars := mux.Vars(r)
			// id := vars["id"]

			if m.RESTStore.Directory != nil {
				time.Sleep(delay)
				w.WriteHeader(m.RESTResponse.DeleteStatusCode)
				for key, val := range m.RESTResponse.Headers {
					w.Header().Set(key, val)
				}

				_ = json.NewEncoder(w).Encode(JSON{
					"message": "not implemented yet!",
				})
			}
		})
	}
}
