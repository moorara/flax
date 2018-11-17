package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	httpMock = HTTPMock{
		Methods:    []string{"POST", "PUT"},
		Path:       "/api/v1/run",
		Queries:    map[string]string{"type": "\\w+"},
		ReqHeaders: map[string]string{"Authorization": "Bearer .*"},
		Delay:      "100ms",
		StatusCode: 200,
		ResHeaders: map[string]string{"Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"},
		Body:       map[string]interface{}{"id": "bec7b2b3-a2a8-4960-8bce-2469c25b370f"},
	}

	restMock = RESTMock{
		BasePath:   "/api/v1/teams",
		ReqHeaders: map[string]string{"Authorization": "Bearer .*"},
		Delay:      "100ms",
		ResHeaders: map[string]string{"Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"},
		Identifier: "_id",
		ListHandle: "data",
		Store: []JSON{
			JSON{"_id": "569358d8-e7a4-4d66-afdc-b343bbfd0c77", "name": "Back-end"},
			JSON{"_id": "4609a5f5-5c55-4d61-9b0d-42b6b35d9013", "name": "Front-end"},
		},
	}
)

func TestHTTPMockWithDefaults(t *testing.T) {
	tests := []struct {
		name         string
		mock         HTTPMock
		expectedMock HTTPMock
	}{
		{
			"DefaultsRequired",
			HTTPMock{},
			HTTPMock{
				Methods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
				Path:       "/",
				Queries:    map[string]string{},
				ReqHeaders: map[string]string{},
				Delay:      "0",
				StatusCode: 200,
				ResHeaders: map[string]string{},
				Body:       nil,
			},
		},
		{
			"DefaultsNotRequired",
			httpMock,
			httpMock,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedMock, tc.mock.withDefaults())
		})
	}
}

func TestHTTPMockHTTPExpectations(t *testing.T) {
	tests := []struct {
		name                     string
		mock                     HTTPMock
		expectedHTTPExpectations []HTTPExpectation
	}{
		{
			"Empty",
			HTTPMock{},
			[]HTTPExpectation{
				HTTPExpectation{
					Methods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
					Path:    "/",
					Queries: []Pair{},
					Headers: []Pair{},
				},
			},
		},
		{
			"NonEmpty",
			httpMock,
			[]HTTPExpectation{
				HTTPExpectation{
					Methods: []string{"POST", "PUT"},
					Path:    "/api/v1/run",
					Queries: []Pair{
						{"type", "{type:\\w+}"},
					},
					Headers: []Pair{
						{"Authorization", "Bearer .*"},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedHTTPExpectations, tc.mock.HTTPExpectations())
		})
	}
}

func TestRESTMockWithDefaults(t *testing.T) {
	tests := []struct {
		name         string
		mock         RESTMock
		expectedMock RESTMock
	}{
		{
			"DefaultsRequired",
			RESTMock{},
			RESTMock{
				BasePath:   "/",
				ReqHeaders: map[string]string{},
				Delay:      "0",
				ResHeaders: map[string]string{},
				Identifier: "",
				ListHandle: "",
				Store:      []JSON{},
			},
		},
		{
			"DefaultsNotRequired",
			restMock,
			restMock,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedMock, tc.mock.withDefaults())
		})
	}
}

func TestRESTMockHTTPExpectations(t *testing.T) {
	tests := []struct {
		name                     string
		mock                     RESTMock
		expectedHTTPExpectations []HTTPExpectation
	}{
		{
			"Empty",
			RESTMock{},
			[]HTTPExpectation{
				HTTPExpectation{
					Methods: []string{"GET"},
					Path:    "/",
					Queries: []Pair{pageNoQuery, pageSizeQuery},
					Headers: []Pair{},
				},

				HTTPExpectation{
					Methods: []string{"POST"},
					Path:    "/",
					Queries: []Pair{pageNoQuery, pageSizeQuery},
					Headers: []Pair{},
				},

				HTTPExpectation{
					Methods: []string{"GET"},
					Path:    "/{id:[-0-9A-Za-z]+}",
					Queries: []Pair{pageNoQuery, pageSizeQuery},
					Headers: []Pair{},
				},

				HTTPExpectation{
					Methods: []string{"PUT"},
					Path:    "/{id:[-0-9A-Za-z]+}",
					Queries: []Pair{pageNoQuery, pageSizeQuery},
					Headers: []Pair{},
				},

				HTTPExpectation{
					Methods: []string{"PATCH"},
					Path:    "/{id:[-0-9A-Za-z]+}",
					Queries: []Pair{pageNoQuery, pageSizeQuery},
					Headers: []Pair{},
				},

				HTTPExpectation{
					Methods: []string{"DELETE"},
					Path:    "/{id:[-0-9A-Za-z]+}",
					Queries: []Pair{pageNoQuery, pageSizeQuery},
					Headers: []Pair{},
				},
			},
		},
		{
			"NonEmpty",
			restMock,
			[]HTTPExpectation{
				HTTPExpectation{
					Methods: []string{"GET"},
					Path:    "/api/v1/teams",
					Queries: []Pair{
						pageNoQuery,
						pageSizeQuery,
						{"name", "{name}"},
					},
					Headers: []Pair{
						{"Authorization", "Bearer .*"},
					},
				},

				HTTPExpectation{
					Methods: []string{"POST"},
					Path:    "/api/v1/teams",
					Queries: []Pair{
						pageNoQuery,
						pageSizeQuery,
						{"name", "{name}"},
					},
					Headers: []Pair{
						{"Authorization", "Bearer .*"},
					},
				},

				HTTPExpectation{
					Methods: []string{"GET"},
					Path:    "/api/v1/teams/{id:[-0-9A-Za-z]+}",
					Queries: []Pair{
						pageNoQuery,
						pageSizeQuery,
						{"name", "{name}"},
					},
					Headers: []Pair{
						{"Authorization", "Bearer .*"},
					},
				},

				HTTPExpectation{
					Methods: []string{"PUT"},
					Path:    "/api/v1/teams/{id:[-0-9A-Za-z]+}",
					Queries: []Pair{
						pageNoQuery,
						pageSizeQuery,
						{"name", "{name}"},
					},
					Headers: []Pair{
						{"Authorization", "Bearer .*"},
					},
				},

				HTTPExpectation{
					Methods: []string{"PATCH"},
					Path:    "/api/v1/teams/{id:[-0-9A-Za-z]+}",
					Queries: []Pair{
						pageNoQuery,
						pageSizeQuery,
						{"name", "{name}"},
					},
					Headers: []Pair{
						{"Authorization", "Bearer .*"},
					},
				},

				HTTPExpectation{
					Methods: []string{"DELETE"},
					Path:    "/api/v1/teams/{id:[-0-9A-Za-z]+}",
					Queries: []Pair{
						pageNoQuery,
						pageSizeQuery,
						{"name", "{name}"},
					},
					Headers: []Pair{
						{"Authorization", "Bearer .*"},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedHTTPExpectations, tc.mock.HTTPExpectations())
		})
	}
}
