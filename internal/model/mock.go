package model

import (
	"fmt"
	"path/filepath"
)

var (
	identifiers = []string{"id", "_id", "Id", "ID"}

	pageNoQuery = Pair{
		Name:  "pageNo",
		Value: "{pageNo:[0-9]+}",
	}

	pageSizeQuery = Pair{
		Name:  "pageSize",
		Value: "{pageSize:[0-9]+}",
	}
)

type (
	// Mock is the type for a mock object
	Mock interface {
		HTTPExpectations() []HTTPExpectation
	}

	// JSON is the type for json objects
	JSON map[string]interface{}

	// HTTPMock is the model for an http mock
	HTTPMock struct {
		Methods    []string          `json:"methods" yaml:"methods"`
		Path       string            `json:"path" yaml:"path"`
		Queries    map[string]string `json:"queries" yaml:"queries"`
		ReqHeaders map[string]string `json:"reqHeaders" yaml:"req_headers"`
		Delay      string            `json:"delay" yaml:"delay"`
		StatusCode int               `json:"statusCode" yaml:"status_code"`
		ResHeaders map[string]string `json:"resHeaders" yaml:"res_headers"`
		Body       interface{}       `json:"body" yaml:"body"`
	}

	// RESTMock is the model for a RESTful resource mock
	RESTMock struct {
		BasePath   string            `json:"basePath" yaml:"base_path"`
		ReqHeaders map[string]string `json:"reqHeaders" yaml:"req_headers"`
		Delay      string            `json:"delay" yaml:"delay"`
		ResHeaders map[string]string `json:"resHeaders" yaml:"res_headers"`
		Identifier string            `json:"identifier" yaml:"identifier"`
		ListHandle string            `json:"listHandle" yaml:"list_handle"`
		Store      []JSON            `json:"store" yaml:"store"`
	}
)

func isID(prop string) bool {
	for _, id := range identifiers {
		if prop == id {
			return true
		}
	}
	return false
}

func (m HTTPMock) withDefaults() HTTPMock {
	if m.Methods == nil || len(m.Methods) == 0 {
		m.Methods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	}

	if m.Path == "" {
		m.Path = "/"
	}

	if m.Queries == nil {
		m.Queries = map[string]string{}
	}

	if m.ReqHeaders == nil {
		m.ReqHeaders = map[string]string{}
	}

	if m.Delay == "" {
		m.Delay = "0"
	}

	if m.StatusCode < 100 || m.StatusCode > 599 {
		m.StatusCode = 200
	}

	if m.ResHeaders == nil {
		m.ResHeaders = map[string]string{}
	}

	return m
}

// HTTPExpectations returns the http expectations for an http mock
func (m HTTPMock) HTTPExpectations() []HTTPExpectation {
	m = m.withDefaults()

	queries := []Pair{}
	for name, regex := range m.Queries {
		queries = append(queries, Pair{
			Name:  name,
			Value: fmt.Sprintf("{%s:%s}", name, regex),
		})
	}

	headers := []Pair{}
	for name, value := range m.ReqHeaders {
		headers = append(headers, Pair{
			Name:  name,
			Value: value,
		})
	}

	return []HTTPExpectation{
		HTTPExpectation{
			Methods: m.Methods,
			Path:    m.Path,
			Queries: queries,
			Headers: headers,
		},
	}
}

func (m RESTMock) withDefaults() RESTMock {
	if m.BasePath == "" {
		m.BasePath = "/"
	}

	if m.ReqHeaders == nil {
		m.ReqHeaders = map[string]string{}
	}

	if m.Delay == "" {
		m.Delay = "0"
	}

	if m.ResHeaders == nil {
		m.ResHeaders = map[string]string{}
	}

	if m.Identifier == "" {
		m.Identifier = ""
	}

	if m.ListHandle == "" {
		m.ListHandle = ""
	}

	if m.Store == nil {
		m.Store = []JSON{}
	}

	return m
}

// HTTPExpectations returns the http expectations for a RESTful resource mock
func (m RESTMock) HTTPExpectations() []HTTPExpectation {
	m = m.withDefaults()
	idPath := fmt.Sprintf("/{id:[-0-9A-Za-z]+}")

	qm := map[string]string{}
	for _, obj := range m.Store {
		for prop := range obj {
			if !isID(prop) {
				qm[prop] = fmt.Sprintf("{%s}", prop)
			}
		}
	}

	// Default query params for pagination
	queries := []Pair{pageNoQuery, pageSizeQuery}
	for name, value := range qm {
		queries = append(queries, Pair{
			Name:  name,
			Value: value,
		})
	}

	headers := []Pair{}
	for name, value := range m.ReqHeaders {
		headers = append(headers, Pair{
			Name:  name,
			Value: value,
		})
	}

	return []HTTPExpectation{
		HTTPExpectation{
			Methods: []string{"GET"},
			Path:    filepath.Join(m.BasePath),
			Queries: queries,
			Headers: headers,
		},

		HTTPExpectation{
			Methods: []string{"POST"},
			Path:    filepath.Join(m.BasePath),
			Queries: queries,
			Headers: headers,
		},

		HTTPExpectation{
			Methods: []string{"GET"},
			Path:    filepath.Join(m.BasePath, idPath),
			Queries: queries,
			Headers: headers,
		},

		HTTPExpectation{
			Methods: []string{"PUT"},
			Path:    filepath.Join(m.BasePath, idPath),
			Queries: queries,
			Headers: headers,
		},

		HTTPExpectation{
			Methods: []string{"PATCH"},
			Path:    filepath.Join(m.BasePath, idPath),
			Queries: queries,
			Headers: headers,
		},

		HTTPExpectation{
			Methods: []string{"DELETE"},
			Path:    filepath.Join(m.BasePath, idPath),
			Queries: queries,
			Headers: headers,
		},
	}
}
