package v1

import (
	"encoding/json"
	"hash/fnv"
	"net/http"
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

// Hash calculates a hash for a rest expectation.
func (e RESTExpect) Hash() uint64 {
	h := fnv.New64a()

	hashString(h, e.BasePath)
	hashStringMap(h, true, e.Headers)

	return h.Sum64()
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

// Hash calculates a hash for a rest mock based on the rest expectation.
func (m RESTMock) Hash() uint64 {
	return m.RESTExpect.Hash()
}

// RegisterRoutes configure routes for a rest mock
func (m RESTMock) RegisterRoutes(router *mux.Router) {
	d, _ := time.ParseDuration(m.Delay)

	// GET /
	{
		path := m.RESTExpect.BasePath
		route := router.Methods("GET").Path(path)
		for header, pattern := range m.RESTExpect.Headers {
			route.HeadersRegexp(header, pattern)
		}

		route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(d)
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

			json.NewEncoder(w).Encode(resp)
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
				time.Sleep(d)
				w.WriteHeader(m.RESTResponse.PostStatusCode)
				for key, val := range m.RESTResponse.Headers {
					w.Header().Set(key, val)
				}

				json.NewEncoder(w).Encode(JSON{
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
				time.Sleep(d)
				w.WriteHeader(m.RESTResponse.GetStatusCode)
				for key, val := range m.RESTResponse.Headers {
					w.Header().Set(key, val)
				}

				json.NewEncoder(w).Encode(JSON{
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
				time.Sleep(d)
				w.WriteHeader(m.RESTResponse.PutStatusCode)
				for key, val := range m.RESTResponse.Headers {
					w.Header().Set(key, val)
				}

				json.NewEncoder(w).Encode(JSON{
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
				time.Sleep(d)
				w.WriteHeader(m.RESTResponse.PatchStatusCode)
				for key, val := range m.RESTResponse.Headers {
					w.Header().Set(key, val)
				}

				json.NewEncoder(w).Encode(JSON{
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
				time.Sleep(d)
				w.WriteHeader(m.RESTResponse.DeleteStatusCode)
				for key, val := range m.RESTResponse.Headers {
					w.Header().Set(key, val)
				}

				json.NewEncoder(w).Encode(JSON{
					"message": "not implemented yet!",
				})
			}
		})
	}
}

// DefaultRESTMock returns a default RESTMock.
func DefaultRESTMock() RESTMock {
	restStore := RESTStore{
		Objects: []JSON{},
	}

	restStore.Index()

	return RESTMock{
		RESTExpect{
			BasePath: "/",
		},
		RESTResponse{
			GetStatusCode:    200,
			PostStatusCode:   201,
			PutStatusCode:    200,
			PatchStatusCode:  200,
			DeleteStatusCode: 204,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			ListKey: "data",
		},
		restStore,
	}
}
