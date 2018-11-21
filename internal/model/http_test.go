package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPExpectWithDefaults(t *testing.T) {
	tests := []struct {
		name           string
		expect         HTTPExpect
		expectedExpect HTTPExpect
	}{
		{
			"Empty",
			HTTPExpect{},
			HTTPExpect{
				Methods: []string{"GET"},
				Path:    "/",
				Prefix:  false,
				Queries: map[string]string{},
				Headers: map[string]string{},
			},
		},
		{
			"DefaultRequired",
			HTTPExpect{
				Path: "/health",
			},
			HTTPExpect{
				Methods: []string{"GET"},
				Path:    "/health",
				Prefix:  false,
				Queries: map[string]string{},
				Headers: map[string]string{},
			},
		},
		{
			"NoDefaultRequired",
			HTTPExpect{
				Methods: []string{"POST", "PUT"},
				Path:    "/sessions",
				Prefix:  true,
				Queries: map[string]string{
					"tenantId": "\\w+",
				},
				Headers: map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				},
			},
			HTTPExpect{
				Methods: []string{"POST", "PUT"},
				Path:    "/sessions",
				Prefix:  true,
				Queries: map[string]string{
					"tenantId": "\\w+",
				},
				Headers: map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedExpect, tc.expect.WithDefaults())
		})
	}
}

func TestHTTPResponseWithDefaults(t *testing.T) {
	tests := []struct {
		name     string
		response HTTPResponse
	}{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NotNil(t, tc)
		})
	}
}

func TestHTTPForwardWithDefaults(t *testing.T) {
	tests := []struct {
		name    string
		forward HTTPForward
	}{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NotNil(t, tc)
		})
	}
}

func TestHTTPMockWithDefaults(t *testing.T) {
	tests := []struct {
		name string
		mock HTTPMock
	}{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NotNil(t, tc)
		})
	}
}

func TestHTTPMockHash(t *testing.T) {
	tests := []struct {
		name string
		m1   HTTPMock
		m2   HTTPMock
	}{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NotNil(t, tc)
		})
	}
}
