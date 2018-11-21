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

func TestHTTPExpectHash(t *testing.T) {
	tests := []struct {
		name          string
		e1            HTTPExpect
		e2            HTTPExpect
		expectedEqual bool
	}{
		{
			"Equal",
			HTTPExpect{
				Methods: []string{"POST", "PUT"},
				Path:    "/v1/sessions",
				Prefix:  true,
				Queries: map[string]string{
					"tenantId": "\\w+",
					"deviceId": "\\w+",
				},
				Headers: map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				},
			},
			HTTPExpect{
				Methods: []string{"POST", "PUT"},
				Path:    "/v1/sessions",
				Prefix:  true,
				Queries: map[string]string{
					"tenantId": "\\w+",
					"deviceId": "\\w+",
				},
				Headers: map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				},
			},
			true,
		},
		{
			"NotEqual",
			HTTPExpect{
				Methods: []string{"POST", "PUT"},
				Path:    "/v1/sessions",
				Prefix:  true,
				Queries: map[string]string{
					"tenantId": "\\w+",
					"deviceId": "\\w+",
				},
				Headers: map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				},
			},
			HTTPExpect{
				Methods: []string{"POST", "PUT"},
				Path:    "/v1/sessions",
				Prefix:  false,
				Queries: map[string]string{
					"tenantId": "\\w+",
					"deviceId": "\\w+",
				},
				Headers: map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				},
			},
			false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expectedEqual {
				assert.Equal(t, tc.e1.Hash(), tc.e2.Hash())
			} else {
				assert.NotEqual(t, tc.e1.Hash(), tc.e2.Hash())
			}
		})
	}
}

func TestHTTPResponseWithDefaults(t *testing.T) {
	tests := []struct {
		name             string
		response         HTTPResponse
		expectedResponse HTTPResponse
	}{
		{
			"Empty",
			HTTPResponse{},
			HTTPResponse{
				Delay:      "0",
				StatusCode: 200,
				Headers:    map[string]string{},
				Body:       nil,
			},
		},
		{
			"DefaultRequired",
			HTTPResponse{
				Delay:      "100ms",
				StatusCode: 201,
			},
			HTTPResponse{
				Delay:      "100ms",
				StatusCode: 201,
				Headers:    map[string]string{},
				Body:       nil,
			},
		},
		{
			"NoDefaultRequired",
			HTTPResponse{
				Delay:      "100ms",
				StatusCode: 201,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: map[string]interface{}{
					"id":    "5da8349a-0707-4064-8fad-74cedb48a8fc",
					"name":  "John Doe",
					"email": "john.doe@example.com",
				},
			},
			HTTPResponse{
				Delay:      "100ms",
				StatusCode: 201,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: map[string]interface{}{
					"id":    "5da8349a-0707-4064-8fad-74cedb48a8fc",
					"name":  "John Doe",
					"email": "john.doe@example.com",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedResponse, tc.response.WithDefaults())
		})
	}
}

func TestHTTPForwardWithDefaults(t *testing.T) {
	tests := []struct {
		name            string
		forward         HTTPForward
		expectedForward HTTPForward
	}{
		{
			"Empty",
			HTTPForward{},
			HTTPForward{
				Delay:   "0",
				To:      "",
				Headers: map[string]string{},
			},
		},
		{
			"DefaultRequired",
			HTTPForward{
				To: "http://service:8080/api",
			},
			HTTPForward{
				Delay:   "0",
				To:      "http://service:8080/api",
				Headers: map[string]string{},
			},
		},
		{
			"NoDefaultRequired",
			HTTPForward{
				Delay: "100ms",
				To:    "http://service:8080/api",
				Headers: map[string]string{
					"Referer": "flax",
				},
			},
			HTTPForward{
				Delay: "100ms",
				To:    "http://service:8080/api",
				Headers: map[string]string{
					"Referer": "flax",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedForward, tc.forward.WithDefaults())
		})
	}
}

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
				HTTPExpect: HTTPExpect{
					Methods: []string{"GET"},
					Path:    "/",
					Prefix:  false,
					Queries: map[string]string{},
					Headers: map[string]string{},
				},
				HTTPResponse: HTTPResponse{
					Delay:      "0",
					StatusCode: 200,
					Headers:    map[string]string{},
					Body:       nil,
				},
				HTTPForward: HTTPForward{
					Delay:   "0",
					To:      "",
					Headers: map[string]string{},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedMock, tc.mock.WithDefaults())
		})
	}
}

func TestHTTPMockHash(t *testing.T) {
	tests := []struct {
		name          string
		m1            HTTPMock
		m2            HTTPMock
		expectedEqual bool
	}{
		{
			"Equal",
			HTTPMock{
				HTTPExpect{
					Methods: []string{"POST", "PUT"},
					Path:    "/v1/sessions",
					Prefix:  true,
					Queries: map[string]string{
						"tenantId": "\\w+",
						"deviceId": "\\w+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPResponse{
					Delay:      "0",
					StatusCode: 200,
					Headers:    map[string]string{},
				},
				HTTPForward{
					Delay:   "0",
					To:      "",
					Headers: map[string]string{},
				},
			},
			HTTPMock{
				HTTPExpect{
					Methods: []string{"POST", "PUT"},
					Path:    "/v1/sessions",
					Prefix:  true,
					Queries: map[string]string{
						"tenantId": "\\w+",
						"deviceId": "\\w+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPResponse{
					Delay:      "100ms",
					StatusCode: 201,
					Headers: map[string]string{
						"Content-Type": "application/json",
					},
				},
				HTTPForward{
					Delay: "100ms",
					To:    "http://service:8080/api",
					Headers: map[string]string{
						"Referer": "flax",
					},
				},
			},
			true,
		},
		{
			"NotEqual",
			HTTPMock{
				HTTPExpect{
					Methods: []string{"POST", "PUT"},
					Path:    "/v1/sessions",
					Prefix:  true,
					Queries: map[string]string{
						"tenantId": "\\w+",
						"deviceId": "\\w+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPResponse{
					Delay:      "0",
					StatusCode: 200,
					Headers:    map[string]string{},
				},
				HTTPForward{
					Delay:   "0",
					To:      "",
					Headers: map[string]string{},
				},
			},
			HTTPMock{
				HTTPExpect{
					Methods: []string{"POST", "PUT"},
					Path:    "/v1/sessions",
					Prefix:  false,
					Queries: map[string]string{
						"tenantId": "\\w+",
						"deviceId": "\\w+",
					},
					Headers: map[string]string{
						"Accept":       "application/json",
						"Content-Type": "application/json",
					},
				},
				HTTPResponse{
					Delay:      "100ms",
					StatusCode: 201,
					Headers: map[string]string{
						"Content-Type": "application/json",
					},
				},
				HTTPForward{
					Delay: "100ms",
					To:    "http://service:8080/api",
					Headers: map[string]string{
						"Referer": "flax",
					},
				},
			},
			false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expectedEqual {
				assert.Equal(t, tc.m1.Hash(), tc.m2.Hash())
			} else {
				assert.NotEqual(t, tc.m1.Hash(), tc.m2.Hash())
			}
		})
	}
}
