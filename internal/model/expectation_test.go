package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanonical(t *testing.T) {
	tests := []struct {
		name                string
		expectation         HTTPExpectation
		expectedExpectation HTTPExpectation
	}{
		{
			"Empty",
			HTTPExpectation{},
			HTTPExpectation{
				Methods: []string{},
				Path:    "/",
				Queries: []Pair{},
				Headers: []Pair{},
			},
		},
		{
			"NonEmpty",
			HTTPExpectation{
				Methods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "TRACE"},
				Path:    "teams",
				Queries: []Pair{
					{"pageNo", "{pageNo:[0-9]+}"},
					{"pageSize", "{pageSize:[0-9]+}"},
				},
				Headers: []Pair{
					{"Content-Type", "application/(text|json)"},
					{"Accept", "application/(text|json)"},
				},
			},
			HTTPExpectation{
				Methods: []string{"DELETE", "GET", "OPTIONS", "PATCH", "POST", "PUT", "TRACE"},
				Path:    "/teams",
				Queries: []Pair{
					{"pageNo", "{pageNo:[0-9]+}"},
					{"pageSize", "{pageSize:[0-9]+}"},
				},
				Headers: []Pair{
					{"Accept", "application/(text|json)"},
					{"Content-Type", "application/(text|json)"},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedExpectation, tc.expectation.canonical())
		})
	}
}

func TestGetHash(t *testing.T) {
	tests := []struct {
		name              string
		httpExpectation01 HTTPExpectation
		httpExpectation02 HTTPExpectation
	}{
		{
			"Empty",
			HTTPExpectation{},
			HTTPExpectation{},
		},
		{
			"Simple",
			HTTPExpectation{
				Methods: []string{"GET"},
				Path:    "/teams",
				Queries: []Pair{
					{"pageNo", "{pageNo:[0-9]+}"},
				},
				Headers: []Pair{
					{"Accept", "application/(text|json)"},
				},
			},
			HTTPExpectation{
				Methods: []string{"GET"},
				Path:    "/teams",
				Queries: []Pair{
					{"pageNo", "{pageNo:[0-9]+}"},
				},
				Headers: []Pair{
					{"Accept", "application/(text|json)"},
				},
			},
		},
		{
			"Complex",
			HTTPExpectation{
				Methods: []string{"POST", "PUT"},
				Path:    "/teams",
				Queries: []Pair{
					{"pageSize", "{pageSize:[0-9]+}"},
					{"pageNo", "{pageNo:[0-9]+}"},
				},
				Headers: []Pair{
					{"Content-Type", "application/(text|json)"},
					{"Accept", "application/(text|json)"},
				},
			},
			HTTPExpectation{
				Methods: []string{"PUT", "POST"},
				Path:    "/teams",
				Queries: []Pair{
					{"pageSize", "{pageSize:[0-9]+}"},
					{"pageNo", "{pageNo:[0-9]+}"},
				},
				Headers: []Pair{
					{"Content-Type", "application/(text|json)"},
					{"Accept", "application/(text|json)"},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.httpExpectation01.GetHash(), tc.httpExpectation02.GetHash())
		})
	}
}
